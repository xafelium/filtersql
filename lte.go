package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// buildLowerThanOrEqual adds the filter.LowerThanOrEqualCondition to the QueryBuilder.
func buildLowerThanOrEqual(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	c, ok := condition.(*filter.LowerThanOrEqualCondition)
	if !ok {
		return fmt.Errorf("condition is no LowerThanOrEqualCondition")
	}
	parameterName := builder.AddArgument(c.Value)
	fieldName, err := mf(c.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s <= %s", fieldName, parameterName))
	return nil
}
