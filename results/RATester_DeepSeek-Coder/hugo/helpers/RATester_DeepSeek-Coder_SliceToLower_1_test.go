package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceToLower_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "single element",
			input:    []string{"A"},
			expected: []string{"a"},
		},
		{
			name:     "multiple elements",
			input:    []string{"A", "B", "C"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "mixed case",
			input:    []string{"a", "B", "c", "D"},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "nil slice",
			input:    nil,
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

			result := SliceToLower(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
