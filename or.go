package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildOr adds the filter.OrCondition to the QueryBuilder.
func buildOr(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	orCondition, ok := condition.(*filter.OrCondition)
	if !ok {
		return fmt.Errorf("condition is no OrCondition")
	}
	if len(orCondition.Conditions) < 2 {
		return fmt.Errorf("OR condition must have at least two conditions")
	}

	for _, c := range orCondition.Conditions {
		subBuilder := builder.SubQueryBuilder()
		err := buildQuery(subBuilder, c, mf)
		if err != nil {
			return err
		}
		if builder.GetSql() != "" {
			builder.SetSql(fmt.Sprintf("%s OR %s", builder.GetSql(), subBuilder.GetSql()))
		} else {
			builder.SetSql(subBuilder.GetSql())
		}
	}

	return nil
}
