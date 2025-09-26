package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStringMap_1(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected map[string]any
		context  *Context
	}

	testCases := []testCase{
		{
			name:     "Test with existing key",
			key:      "existingKey",
			expected: map[string]any{"key1": "value1", "key2": "value2"},
			context: &Context{
				Keys: map[string]any{
					"existingKey": map[string]any{"key1": "value1", "key2": "value2"},
				},
			},
		},
		{
			name:     "Test with non-existing key",
			key:      "nonExistingKey",
			expected: nil,
			context: &Context{
				Keys: map[string]any{
					"existingKey": map[string]any{"key1": "value1", "key2": "value2"},
				},
			},
		},
		{
			name:     "Test with nil value",
			key:      "nilKey",
			expected: nil,
			context: &Context{
				Keys: map[string]any{
					"nilKey": nil,
				},
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

			result := tc.context.GetStringMap(tc.key)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
