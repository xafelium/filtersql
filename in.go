package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
	"reflect"
	"strings"
)

// buildInCondition adds the filter.InCondition to the QueryBuilder.
func buildInCondition(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	c, ok := condition.(*filter.InCondition)
	if !ok {
		return fmt.Errorf("condition is no InCondition")
	}
	var parameterNames []string
	v := reflect.ValueOf(c.Value)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			val := v.Index(i).Interface()
			parameterNames = append(parameterNames, builder.AddArgument(val))
		}
	default:
		parameterNames = append(parameterNames, builder.AddArgument(c.Value))
	}
	fieldName, err := mf(c.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s IN (%s)", fieldName, strings.Join(parameterNames, ", ")))
	return nil
}
