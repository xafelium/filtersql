package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

func buildNot(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	notCondition, ok := condition.(*filter.NotCondition)
	if !ok {
		return fmt.Errorf("condition is no NotCondition")
	}
	subBuilder := builder.SubQueryBuilder()
	err := buildQuery(subBuilder, notCondition.Condition, mf)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("NOT (%s)", subBuilder.GetSql()))

	return nil
}
