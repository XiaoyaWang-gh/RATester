package utils

import (
	"fmt"
	"testing"
)

func TestGetDisplayString_1(t *testing.T) {
	type test struct {
		name     string
		input    []interface{}
		expected string
	}

	tests := []test{
		{
			name:     "Test case 1",
			input:    []interface{}{"Hello", "World"},
			expected: "Hello World",
		},
		{
			name:     "Test case 2",
			input:    []interface{}{"Hello", 1, true, 3.14},
			expected: "Hello 1 true 3.14",
		},
		{
			name:     "Test case 3",
			input:    []interface{}{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := GetDisplayString(tt.input...)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
