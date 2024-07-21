package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildLowerThan adds the filter.LowerThanCondition to the QueryBuilder.
func buildLowerThan(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	c, ok := condition.(*filter.LowerThanCondition)
	if !ok {
		return fmt.Errorf("condition is no LowerThanCondition")
	}
	parameterName := builder.AddArgument(c.Value)
	fieldName, err := mf(c.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s < %s", fieldName, parameterName))
	return nil
}
