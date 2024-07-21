package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildGreaterThanOrEqual adds the filter.GreaterThanOrEqualCondition to the QueryBuilder.
func buildGreaterThanOrEqual(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	c, ok := condition.(*filter.GreaterThanOrEqualCondition)
	if !ok {
		return fmt.Errorf("condition is no GreaterThanOrEqualCondition")
	}
	parameterName := builder.AddArgument(c.Value)
	fieldName, err := mf(c.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s >= %s", fieldName, parameterName))
	return nil
}
