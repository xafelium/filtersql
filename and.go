package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildAnd adds the "AND" condition to the QueryBuilder.
func buildAnd(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	andCondition, ok := condition.(*filter.AndCondition)
	if !ok {
		return fmt.Errorf("condition is no AndCondition")
	}
	if len(andCondition.Conditions) < 2 {
		return fmt.Errorf("AND condition must have at least two conditions")
	}

	for _, c := range andCondition.Conditions {
		subBuilder := builder.SubQueryBuilder()
		err := buildQuery(subBuilder, c, mf)
		if err != nil {
			return err
		}
		if builder.GetSql() != "" {
			builder.SetSql(fmt.Sprintf("%s AND %s", builder.GetSql(), subBuilder.GetSql()))
		} else {
			builder.SetSql(subBuilder.GetSql())
		}
	}

	return nil
}
