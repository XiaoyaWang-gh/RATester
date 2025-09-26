package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakePathsSanitized_1(t *testing.T) {
	type test struct {
		name     string
		input    []string
		expected []string
	}

	tests := []test{
		{
			name:     "empty",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "single",
			input:    []string{"/path/to/file"},
			expected: []string{"/path/to/file"},
		},
		{
			name:     "multiple",
			input:    []string{"/path/to/file", "/another/path"},
			expected: []string{"/path/to/file", "/another/path"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ps := &PathSpec{}
			ps.MakePathsSanitized(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, test.input)
			}
		})
	}
}
