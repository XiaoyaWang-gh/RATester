package transform

import (
	"fmt"
	"reflect"
	"testing"
)

func TestApplyMarshalTypes_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "Test case 1",
			input: map[string]any{
				"key1": float64(123.0),
				"key2": map[string]any{
					"key3": float64(456.0),
				},
			},
			expected: map[string]any{
				"key1": int64(123),
				"key2": map[string]any{
					"key3": int64(456),
				},
			},
		},
		{
			name: "Test case 2",
			input: map[string]any{
				"key1": float64(789.0),
				"key2": map[string]any{
					"key3": float64(0.0),
				},
			},
			expected: map[string]any{
				"key1": int64(789),
				"key2": map[string]any{
					"key3": int64(0),
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

			applyMarshalTypes(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.input)
			}
		})
	}
}
