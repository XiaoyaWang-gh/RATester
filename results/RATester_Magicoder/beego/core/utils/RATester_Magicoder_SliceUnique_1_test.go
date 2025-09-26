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
			name:     "Test case 1",
			input:    []interface{}{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
			expected: []interface{}{1, 2, 3, 4, 5},
		},
		{
			name:     "Test case 2",
			input:    []interface{}{"a", "b", "c", "a", "b", "c"},
			expected: []interface{}{"a", "b", "c"},
		},
		{
			name:     "Test case 3",
			input:    []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "Test case 4",
			input:    []interface{}{1, "a", 2, "b", 3, "c"},
			expected: []interface{}{1, "a", 2, "b", 3, "c"},
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
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
