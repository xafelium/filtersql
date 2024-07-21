package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildGreaterThan adds the filter.GreaterThanCondition to the QueryBuilder.
func buildGreaterThan(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	c, ok := condition.(*filter.GreaterThanCondition)
	if !ok {
		return fmt.Errorf("condition is no GreaterThanCondition")
	}
	parameterName := builder.AddArgument(c.Value)
	fieldName, err := mf(c.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s > %s", fieldName, parameterName))
	return nil
}
