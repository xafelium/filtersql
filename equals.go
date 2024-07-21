package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildEquals adds the filter.EqualsCondition to the QueryBuilder.
func buildEquals(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	equalsCondition, ok := condition.(*filter.EqualsCondition)
	if !ok {
		return fmt.Errorf("condition is no EqualsCondition")
	}
	parameterName := builder.AddArgument(equalsCondition.Value)
	fieldName, err := mf(equalsCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s = %s", fieldName, parameterName))
	return nil
}
