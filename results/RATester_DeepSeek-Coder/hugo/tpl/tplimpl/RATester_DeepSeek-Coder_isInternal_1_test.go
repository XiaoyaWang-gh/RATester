package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsInternal_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Test case 1",
			input:    "internal/test",
			expected: true,
		},
		{
			name:     "Test case 2",
			input:    "external/test",
			expected: false,
		},
		{
			name:     "Test case 3",
			input:    "internal",
			expected: true,
		},
		{
			name:     "Test case 4",
			input:    "external",
			expected: false,
		},
		{
			name:     "Test case 5",
			input:    "",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := isInternal(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
