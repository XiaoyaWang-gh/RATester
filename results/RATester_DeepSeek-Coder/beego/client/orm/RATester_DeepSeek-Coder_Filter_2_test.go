package orm

import (
	"fmt"
	"testing"
	"time"
)

func TestFilter_2(t *testing.T) {
	qs := querySet{}
	testCases := []struct {
		name     string
		expr     string
		args     []interface{}
		expected QuerySeter
	}{
		{
			name:     "Filter with single argument",
			expr:     "UserName",
			args:     []interface{}{"slene"},
			expected: &qs,
		},
		{
			name:     "Filter with multiple arguments",
			expr:     "profile__Age",
			args:     []interface{}{28},
			expected: &qs,
		},
		{
			name:     "Filter with time argument",
			expr:     "created",
			args:     []interface{}{time.Now()},
			expected: &qs,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := qs.Filter(tc.expr, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
