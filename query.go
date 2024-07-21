package filtersql

import (
	"fmt"
	"github.com/xafelium/filter"
)

// FieldMapperFunc is a function to map domain object field names to database table columns.
type FieldMapperFunc func(fieldName string) (string, error)

var (
	conditionBuilders = make(map[string]func(builder *QueryBuilder, c filter.Condition, mf FieldMapperFunc) error)
)

// Query contains the data of a SQL query. It can be used with pgx.
type Query struct {
	Sql  string
	Args []interface{}
}

func init() {
	conditionBuilders[filter.WhereConditionType] = buildWhere
	conditionBuilders[filter.GroupConditionType] = buildGroup
	conditionBuilders[filter.AndConditionType] = buildAnd
	conditionBuilders[filter.OrConditionType] = buildOr
	conditionBuilders[filter.EqualsConditionType] = buildEquals
	conditionBuilders[filter.ContainsConditionType] = buildContains
	conditionBuilders[filter.ArrayContainsConditionType] = buildArrayContains
	conditionBuilders[filter.GreaterThanConditionType] = buildGreaterThan
	conditionBuilders[filter.GreaterThanOrEqualConditionType] = buildGreaterThanOrEqual
	conditionBuilders[filter.LowerThanConditionType] = buildLowerThan
	conditionBuilders[filter.LowerThanOrEqualConditionType] = buildLowerThanOrEqual
	conditionBuilders[filter.OverlapsConditionType] = buildOverlaps
	conditionBuilders[filter.InConditionType] = buildInCondition
	conditionBuilders[filter.IsNilConditionType] = buildIsNil
	conditionBuilders[filter.NotConditionType] = buildNot
	conditionBuilders[filter.NotEqualsConditionType] = buildNotEquals
	conditionBuilders[filter.NotNilConditionType] = buildNotNil
	conditionBuilders[filter.NotRegexConditionType] = buildNotRegex
	conditionBuilders[filter.RegexConditionType] = buildRegex
}

// BuildQuery creates a new Query for the filter.Condition using the FieldMapperFunc.
func BuildQuery(c filter.Condition, mf FieldMapperFunc) (*Query, error) {
	builder := QueryBuilder{}
	err := buildQuery(&builder, c, mf)
	if err != nil {
		return nil, err
	}
	return builder.Build(), nil
}

// buildQuery creates a new Query for the filter.Condition by using the provided QueryBuilder and FieldMapperFunc.
func buildQuery(builder *QueryBuilder, condition filter.Condition, mf FieldMapperFunc) error {
	if condition == nil {
		return nil
	}
	buildFunc, ok := conditionBuilders[condition.Type()]
	if !ok {
		return fmt.Errorf(fmt.Sprintf("unknown condition: %s", condition.Type()))
	}
	return buildFunc(builder, condition, mf)
}
