package filtersql_test

import (
	"fmt"
	"github.com/xafelium/filtersql"

	"github.com/stretchr/testify/require"
	"github.com/xafelium/filter"
	"testing"
)

func TestBuildQuery(t *testing.T) {
	tests := []struct {
		name       string
		filter     filter.Condition
		expected   string
		args       []interface{}
		err        error
		mapperFunc filtersql.FieldMapperFunc
	}{
		{
			name:     "empty where",
			filter:   filter.Where(nil),
			expected: "",
		},
		{
			name: "where with equals",
			filter: filter.Where(
				filter.Equals("id", 1234),
			),
			expected: "WHERE id = $1",
			args:     []interface{}{1234},
		},
		{
			name: "where with not",
			filter: filter.Where(
				filter.Not(
					filter.GreaterThanOrEqual("keys", 123),
				),
			),
			expected: "WHERE NOT (keys >= $1)",
			args:     []any{123},
		},
		{
			name: "where with not equals",
			filter: filter.Where(
				filter.NotEquals("name", "Mustermann"),
			),
			expected: "WHERE name != $1",
			args:     []any{"Mustermann"},
		},
		{
			name: "where with regex",
			filter: filter.Where(
				filter.Regex("name", "4$"),
			),
			expected: "WHERE name ~ $1",
			args:     []interface{}{"4$"},
		},
		{
			name: "where with not regex",
			filter: filter.Where(
				filter.NotRegex("label", "^abc"),
			),
			expected: "WHERE label !~ $1",
			args:     []interface{}{"^abc"},
		},
		{
			name: "where with equals or equals",
			filter: filter.Where(
				filter.Or(
					filter.Equals("label", "abc"),
					filter.Equals("label", "def"),
				),
			),
			expected: "WHERE label = $1 OR label = $2",
			args:     []interface{}{"abc", "def"},
		},
		{
			name: "multiple ors",
			filter: filter.Where(
				filter.Or(
					filter.Equals("label", "abc"),
					filter.Equals("label", "def"),
					filter.Equals("label", "ghi"),
					filter.Equals("label", "jkl"),
				),
			),
			expected: "WHERE label = $1 OR label = $2 OR label = $3 OR label = $4",
			args:     []interface{}{"abc", "def", "ghi", "jkl"},
		},
		{
			name: "empty or",
			filter: filter.Where(
				filter.Or(),
			),
			err: fmt.Errorf("OR condition must have at least two conditions"),
		},
		{
			name: "singe or argument",
			filter: filter.Or(
				filter.Equals("a", "a"),
			),
			err: fmt.Errorf("OR condition must have at least two conditions"),
		},
		{
			name:   "empty and",
			filter: filter.Where(filter.And()),
			err:    fmt.Errorf("AND condition must have at least two conditions"),
		},
		{
			name:   "single and argument",
			filter: filter.Where(filter.And(filter.Equals("a", "a"))),
			err:    fmt.Errorf("AND condition must have at least two conditions"),
		},
		{
			name: "multiple ands",
			filter: filter.Where(
				filter.And(
					filter.Equals("label", "abc"),
					filter.Equals("label", "def"),
					filter.Equals("label", "ghi"),
					filter.Equals("label", "jkl"),
				),
			),
			expected: "WHERE label = $1 AND label = $2 AND label = $3 AND label = $4",
			args:     []interface{}{"abc", "def", "ghi", "jkl"},
		},
		{
			name: "grouping",
			filter: filter.Where(
				filter.And(
					filter.Equals("a", 2),
					filter.Group(
						filter.Or(
							filter.Equals("b", "b"),
							filter.Equals("c", "c"),
							filter.GreaterThan("d", "e"),
							filter.GreaterThanOrEqual("f", "g"),
							filter.LowerThan("h", "i"),
							filter.LowerThanOrEqual("j", "k"),
						),
					),
				),
			),
			expected: "WHERE a = $1 AND (b = $2 OR c = $3 OR d > $4 OR f >= $5 OR h < $6 OR j <= $7)",
			args:     []interface{}{2, "b", "c", "e", "g", "i", "k"},
		},
		{
			name: "contains",
			filter: filter.Where(
				filter.Contains("label", "abc"),
			),
			expected: "WHERE label ILIKE $1",
			args:     []interface{}{"%abc%"},
		},
		{
			name:   "contains with invalid field name",
			filter: filter.Contains("notExistingFieldName", "abc"),
			err:    fmt.Errorf("field name error"),
			mapperFunc: func(fieldName string) (string, error) {
				if fieldName == "notExistingFieldName" {
					return "", fmt.Errorf("field name error")
				}
				return "", fmt.Errorf("should not be reached")
			},
		},
		{
			name:   "equals with invalid field name",
			filter: filter.Equals("notExistingFieldName", "abc"),
			err:    fmt.Errorf("field name error"),
			mapperFunc: func(fieldName string) (string, error) {
				if fieldName == "notExistingFieldName" {
					return "", fmt.Errorf("field name error")
				}
				return "", fmt.Errorf("should not be reached")
			},
		},
		{
			name: "in with slice",
			filter: filter.Where(
				filter.In("id", []int{1, 2}),
			),
			expected: "WHERE id IN ($1, $2)",
			args:     []interface{}{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test := test
			mapperFunc := test.mapperFunc
			if mapperFunc == nil {
				mapperFunc = func(fieldName string) (string, error) {
					return fieldName, nil
				}
			}

			q, err := filtersql.BuildQuery(test.filter, mapperFunc)

			if test.err != nil {
				require.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, q.Sql)
				require.Equal(t, test.args, q.Args)
			}
		})
	}
}
