package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCleanConfigStringMapString_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]string
		expected map[string]string
	}{
		{
			name:     "empty map",
			input:    map[string]string{},
			expected: map[string]string{},
		},
		{
			name:     "map without MergeStrategyKey",
			input:    map[string]string{"key1": "value1", "key2": "value2"},
			expected: map[string]string{"key1": "value1", "key2": "value2"},
		},
		{
			name:     "map with MergeStrategyKey",
			input:    map[string]string{MergeStrategyKey: "value1", "key2": "value2"},
			expected: map[string]string{"key2": "value2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := CleanConfigStringMapString(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
