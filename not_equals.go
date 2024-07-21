package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

func buildNotEquals(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	notEqualsCondition, ok := condition.(*filter.NotEqualsCondition)
	if !ok {
		return fmt.Errorf("condition is no NotEqualsCondition")
	}
	parameterName := builder.AddArgument(notEqualsCondition.Value)
	fieldName, err := mf(notEqualsCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s != %s", fieldName, parameterName))
	return nil
}
