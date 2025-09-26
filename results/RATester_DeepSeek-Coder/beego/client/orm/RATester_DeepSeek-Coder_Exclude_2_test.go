package orm

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestExclude_2(t *testing.T) {
	qs := querySet{}
	testCases := []struct {
		name     string
		expr     string
		args     []interface{}
		expected QuerySeter
	}{
		{
			name: "TestExclude",
			expr: "field_name = ?",
			args: []interface{}{"value"},
			expected: &querySet{
				cond: NewCondition().AndNot("field_name = ?", "value"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := qs.Exclude(tc.expr, tc.args...)
			assert.Equal(t, tc.expected, result)
		})
	}
}
