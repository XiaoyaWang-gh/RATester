package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqueNonEmptyStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    []string
		expected []string
	}

	tests := []test{
		{
			input:    []string{"hello", "world", "hello", "", "world"},
			expected: []string{"hello", "world"},
		},
		{
			input:    []string{"", "", "", "", ""},
			expected: []string{},
		},
		{
			input:    []string{"one", "two", "three", "four", "five"},
			expected: []string{"one", "two", "three", "four", "five"},
		},
		{
			input:    []string{"same", "same", "same", "same", "same"},
			expected: []string{"same"},
		},
	}

	for _, test := range tests {
		result := uniqueNonEmptyStrings(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}
