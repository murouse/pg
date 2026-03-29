package pg

import sq "github.com/Masterminds/squirrel"

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
