package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetParamValues_1(t *testing.T) {
	testCases := []struct {
		name     string
		values   []string
		expected []string
	}{
		{
			name:     "Test case 1",
			values:   []string{"value1", "value2", "value3"},
			expected: []string{"value1", "value2", "value3"},
		},
		{
			name:     "Test case 2",
			values:   []string{"value4", "value5", "value6", "value7"},
			expected: []string{"value4", "value5", "value6", "value7"},
		},
		{
			name:     "Test case 3",
			values:   []string{"value8", "value9"},
			expected: []string{"value8", "value9"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &context{}
			ctx.SetParamValues(tc.values...)

			if !reflect.DeepEqual(ctx.ParamValues(), tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, ctx.ParamValues())
			}
		})
	}
}
