package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildGroup adds the filter.GroupCondition to the QueryBuilder.
func buildGroup(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	groupCondition, ok := condition.(*filter.GroupCondition)
	if !ok {
		return fmt.Errorf("condition is no GroupCondition")
	}
	subBuilder := builder.SubQueryBuilder()
	err := buildQuery(subBuilder, groupCondition.Condition, mf)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("(%s)", subBuilder.sql))
	return nil
}
