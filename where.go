package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildWhere adds the filter.WhereCondition to the QueryBuilder.
func buildWhere(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	if builder.GetSql() != "" {
		return fmt.Errorf(
			"the WhereCondition must be called as first element but the current SQL-String already contains: \"%s\"",
			builder.GetSql())
	}
	whereCondition, ok := condition.(*filter.WhereCondition)
	if !ok {
		return fmt.Errorf("condition is no WhereCondition")
	}
	err := buildQuery(builder, whereCondition.Condition, mf)
	if err != nil {
		return err
	}
	if whereCondition.Condition == nil {
		builder.SetSql("")
		return nil
	}
	builder.SetSql(fmt.Sprintf("WHERE %s", builder.sql))
	return nil
}
