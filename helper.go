package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // register postgres dialect for goqu
	goquexp "github.com/doug-martin/goqu/v9/exp"
)

// Sq returns squirrel statement builder configured for PostgreSQL ($ placeholders).
func Sq() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

// SqlQuery is a simple Sqlizer implementation for raw queries.
type SqlQuery struct {
	query string
	args  []any
}

// ToSql returns raw query and arguments.
func (pq *SqlQuery) ToSql() (string, []any, error) {
	return pq.query, pq.args, nil
}

// Sql creates a new raw SQL query wrapper.
func Sql(query string, args ...interface{}) *SqlQuery {
	return &SqlQuery{
		query: query,
		args:  args,
	}
}

// GoQuWrapper is an adapter that makes goqu expressions compatible with squirrel.Sqlizer.
//
// Reason:
// - goqu uses ToSQL()
// - squirrel expects ToSql()
// This bridges the API mismatch between the two libraries.
type GoQuWrapper struct {
	goquexp.SQLExpression
}

// ToSql adapts goqu's ToSQL method to match squirrel's Sqlizer interface.
func (w *GoQuWrapper) ToSql() (string, []any, error) {
	return w.ToSQL()
}

// GoQu wraps a goqu expression into a Sqlizer-compatible type.
//
// Usage:
//
//	c.Exec(ctx, GoQu(GoQuDialect().Select(...)))
func GoQu(expr goquexp.SQLExpression) *GoQuWrapper {
	return &GoQuWrapper{expr}
}

// GoQuDialect returns a PostgreSQL dialect for goqu.
// Extracted into a helper for consistency and convenience.
func GoQuDialect() goqu.DialectWrapper {
	return goqu.Dialect("postgres")
}
