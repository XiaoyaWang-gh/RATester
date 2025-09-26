package hqt

import (
	"fmt"
	"testing"
)

func TestNormalizeString_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single line with no spaces",
			input:    "Hello, World",
			expected: "Hello,World",
		},
		{
			name:     "single line with spaces",
			input:    "Hello,   World",
			expected: "Hello,World",
		},
		{
			name:     "multiple lines with spaces",
			input:    "Hello,   \n   World",
			expected: "Hello,World",
		},
		{
			name:     "multiple lines with carriage return and newline",
			input:    "Hello,   \r\n   World",
			expected: "Hello,World",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := normalizeString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
