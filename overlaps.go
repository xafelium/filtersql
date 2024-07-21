package filtersql

import (
	fmt "fmt"
	"github.com/xafelium/filter"
)

// buildOverlaps adds the filter.OverlapsCondition to the QueryBuilder.
func buildOverlaps(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	c, ok := condition.(*filter.OverlapsCondition)
	if !ok {
		return fmt.Errorf("condition is no OverlapsCondition")
	}
	parameterName := builder.AddArgument(c.Value)
	fieldName, err := mf(c.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s && %s", fieldName, parameterName))
	return nil
}
