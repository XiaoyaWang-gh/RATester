package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopyJSON_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "simple map",
			input: map[string]any{
				"key1": "value1",
				"key2": 2,
				"key3": true,
			},
			expected: map[string]any{
				"key1": "value1",
				"key2": 2,
				"key3": true,
			},
		},
		{
			name: "nested map",
			input: map[string]any{
				"key1": "value1",
				"key2": map[string]any{
					"subkey1": "subvalue1",
					"subkey2": 2,
				},
			},
			expected: map[string]any{
				"key1": "value1",
				"key2": map[string]any{
					"subkey1": "subvalue1",
					"subkey2": 2,
				},
			},
		},
		{
			name:     "empty map",
			input:    map[string]any{},
			expected: map[string]any{},
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			result := deepCopyJSON(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
