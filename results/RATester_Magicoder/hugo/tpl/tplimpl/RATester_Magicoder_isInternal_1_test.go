package tplimpl

import (
	"fmt"
	"testing"
)

func TestisInternal_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Test case 1: Valid internal path",
			input:    "internal/path",
			expected: true,
		},
		{
			name:     "Test case 2: Valid internal path with trailing slash",
			input:    "internal/path/",
			expected: true,
		},
		{
			name:     "Test case 3: Valid internal path with leading slash",
			input:    "/internal/path",
			expected: true,
		},
		{
			name:     "Test case 4: Valid internal path with leading and trailing slash",
			input:    "/internal/path/",
			expected: true,
		},
		{
			name:     "Test case 5: Invalid internal path",
			input:    "external/path",
			expected: false,
		},
		{
			name:     "Test case 6: Invalid internal path with trailing slash",
			input:    "external/path/",
			expected: false,
		},
		{
			name:     "Test case 7: Invalid internal path with leading slash",
			input:    "/external/path",
			expected: false,
		},
		{
			name:     "Test case 8: Invalid internal path with leading and trailing slash",
			input:    "/external/path/",
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
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
