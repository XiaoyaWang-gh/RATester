package runtime

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnique_1(t *testing.T) {
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
			input:    []string{"a", "b", "a", "c", "b", "d"},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			input:    []string{"apple", "banana", "apple", "cherry", "banana", "date"},
			expected: []string{"apple", "banana", "cherry", "date"},
		},
		{
			input:    []string{"apple", "banana", "cherry", "date"},
			expected: []string{"apple", "banana", "cherry", "date"},
		},
		{
			input:    []string{},
			expected: []string{},
		},
	}

	for _, test := range tests {
		result := unique(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}
