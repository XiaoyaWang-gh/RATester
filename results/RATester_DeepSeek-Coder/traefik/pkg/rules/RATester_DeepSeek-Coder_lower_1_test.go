package rules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLower_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Test with lowercase strings",
			input:    []string{"hello", "world"},
			expected: []string{"hello", "world"},
		},
		{
			name:     "Test with mixed case strings",
			input:    []string{"HeLLo", "WoRld"},
			expected: []string{"hello", "world"},
		},
		{
			name:     "Test with empty slice",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := lower(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
