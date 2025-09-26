package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestAggregate_1(t *testing.T) {
	t.Parallel()

	d := &DoNothingQuerySetter{}

	testCases := []struct {
		name     string
		input    string
		expected orm.QuerySeter
	}{
		{
			name:     "Test Aggregate",
			input:    "SUM(age)",
			expected: d,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := d.Aggregate(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
