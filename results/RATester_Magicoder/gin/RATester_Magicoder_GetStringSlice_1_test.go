package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStringSlice_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		value    any
		expected []string
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			value:    []string{"value1", "value2"},
			expected: []string{"value1", "value2"},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			value:    nil,
			expected: nil,
		},
		{
			name:     "Test case 3",
			key:      "key3",
			value:    "invalid",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{
				Keys: make(map[string]any),
			}
			ctx.Keys[tc.key] = tc.value

			result := ctx.GetStringSlice(tc.key)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
