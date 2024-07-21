package filtersql

import (
	"fmt"
	"strings"
)

// QueryBuilder builds sql statements for usage with pgx.
type QueryBuilder struct {
	sql    string
	args   []interface{}
	parent *QueryBuilder
}

// SubQueryBuilder returns a new QueryBuilder containing the current builder instance.
func (b *QueryBuilder) SubQueryBuilder() *QueryBuilder {
	return &QueryBuilder{parent: b}
}

// GetSql returns the SQL string.
func (b *QueryBuilder) GetSql() string {
	return b.sql
}

// SetSql sets the SQL string.
func (b *QueryBuilder) SetSql(sql string) {
	b.sql = sql
}

// AddArgument adds the value as a positional argument to be used in SQL statements.
func (b *QueryBuilder) AddArgument(value interface{}) string {
	if b.parent != nil {
		return b.parent.AddArgument(value)
	}
	b.args = append(b.args, value)
	return fmt.Sprintf("###%d###", len(b.args))
}

// Build returns a new Query object for the current QueryBuilder object.
func (b *QueryBuilder) Build() *Query {
	for i := range b.args {
		b.sql = strings.Replace(b.sql, fmt.Sprintf("###%d###", i+1), fmt.Sprintf("$%d", i+1), -1)
	}
	return &Query{
		Sql:  b.sql,
		Args: b.args,
	}
}
