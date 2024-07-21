package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildContains adds the filter.ContainsCondition to the QueryBuilder.
func buildContains(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	containsCondition, ok := condition.(*filter.ContainsCondition)
	if !ok {
		return fmt.Errorf("condition is no EqualsCondition")
	}
	parameterName := builder.AddArgument(fmt.Sprintf("%%%s%%", containsCondition.Value))
	fieldName, err := mf(containsCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s ILIKE %s", fieldName, parameterName))
	return nil
}
