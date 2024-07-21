package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildArrayContains adds the filter.ArrayContainsCondition to the QueryBuilder.
func buildArrayContains(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	containsCondition, ok := condition.(*filter.ArrayContainsCondition)
	if !ok {
		return fmt.Errorf("condition is no ArrayContainsCondition")
	}
	parameterName := builder.AddArgument(containsCondition.Value)
	fieldName, err := mf(containsCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s=ANY(%s)", parameterName, fieldName))
	return nil
}
