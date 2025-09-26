package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttemplateVariants_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Test case 1",
			input:    "test",
			expected: []string{"test"},
		},
		{
			name:     "Test case 2",
			input:    "test2",
			expected: []string{"test2"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := templateVariants(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
