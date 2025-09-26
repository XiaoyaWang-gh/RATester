package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStringMapString_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected map[string]string
		context  *Context
	}

	testCases := []testCase{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: map[string]string{"key1": "value1"},
			context: &Context{
				Keys: map[string]any{"key1": map[string]string{"key1": "value1"}},
			},
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: map[string]string{"key2": "value2"},
			context: &Context{
				Keys: map[string]any{"key2": map[string]string{"key2": "value2"}},
			},
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: map[string]string{"key3": "value3"},
			context: &Context{
				Keys: map[string]any{"key3": map[string]string{"key3": "value3"}},
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

			result := tc.context.GetStringMapString(tc.key)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
