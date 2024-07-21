package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

func buildNotRegex(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	notRegexCondition, ok := condition.(*filter.NotRegexCondition)
	if !ok {
		return fmt.Errorf("condition is no NotRegexCondition")
	}
	parameterName := builder.AddArgument(notRegexCondition.Expression)
	fieldName, err := mf(notRegexCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s !~ %s", fieldName, parameterName))
	return nil
}
