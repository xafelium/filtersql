package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

func buildRegex(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	regexCondition, ok := condition.(*filter.RegexCondition)
	if !ok {
		return fmt.Errorf("condition is no RegexCondition")
	}
	parameterName := builder.AddArgument(regexCondition.Expression)
	fieldName, err := mf(regexCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s ~ %s", fieldName, parameterName))
	return nil
}
