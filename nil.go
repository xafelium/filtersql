package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

func buildIsNil(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	isNilCondition, ok := condition.(*filter.IsNilCondition)
	if !ok {
		return fmt.Errorf("condition is no IsNilCondition")
	}
	fieldName, err := mf(isNilCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s IS NULL", fieldName))
	return nil
}

func buildNotNil(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	notNilCondition, ok := condition.(*filter.NotNilCondition)
	if !ok {
		return fmt.Errorf("condition is no NotNilCondition")
	}
	fieldName, err := mf(notNilCondition.Field)
	if err != nil {
		return err
	}
	builder.SetSql(fmt.Sprintf("%s IS NOT NULL", fieldName))
	return nil
}
