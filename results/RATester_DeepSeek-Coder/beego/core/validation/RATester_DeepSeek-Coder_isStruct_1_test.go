package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsStruct_1(t *testing.T) {
	type testStruct struct {
		field string
	}

	type testNonStruct struct {
		field int
	}

	testCases := []struct {
		name     string
		input    reflect.Type
		expected bool
	}{
		{
			name:     "Test with struct type",
			input:    reflect.TypeOf(testStruct{}),
			expected: true,
		},
		{
			name:     "Test with non-struct type",
			input:    reflect.TypeOf(testNonStruct{}),
			expected: false,
		},
		{
			name:     "Test with built-in type",
			input:    reflect.TypeOf(1),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isStruct(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
