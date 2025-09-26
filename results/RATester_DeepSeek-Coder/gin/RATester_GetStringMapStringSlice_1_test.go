package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStringMapStringSlice_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected map[string][]string
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: map[string][]string{"key1": {"value1", "value2"}},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: map[string][]string{"key2": {"value3", "value4"}},
		},
		{
			name:     "Test case 3",
			key:      "key3",
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

			c := &Context{
				Keys: map[string]any{
					tc.key: tc.expected,
				},
			}

			result := c.GetStringMapStringSlice(tc.key)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
