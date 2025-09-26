package provider

import (
	"fmt"
	"testing"
)

func TestNormalize_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "normal case",
			input:    "Hello World",
			expected: "Hello-World",
		},
		{
			name:     "special characters",
			input:    "Hello$World",
			expected: "Hello-World",
		},
		{
			name:     "multiple spaces",
			input:    "Hello   World",
			expected: "Hello-World",
		},
		{
			name:     "leading and trailing spaces",
			input:    "   Hello World   ",
			expected: "Hello-World",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := Normalize(test.input)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
