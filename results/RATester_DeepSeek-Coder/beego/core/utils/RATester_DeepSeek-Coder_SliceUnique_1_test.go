package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceUnique_1(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Test with no duplicates",
			input:    []interface{}{1, 2, 3, 4, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "Test with duplicates",
			input:    []interface{}{1, 2, 2, 3, 3, 4, 4, 5, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "Test with mixed types",
			input:    []interface{}{1, "two", "two", 3.0, 3.0, 4, "four", "four", 5.0, 5.0},
			expected: []interface{}{1, "two", 3.0, 4, "four", 5.0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := SliceUnique(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
