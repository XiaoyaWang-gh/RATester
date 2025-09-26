package page

import (
	"fmt"
	"testing"
)

func TesttrimSpace_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "TrimSpace_EmptyString",
			input:    "",
			expected: "",
		},
		{
			name:     "TrimSpace_StringWithLeadingAndTrailingSpaces",
			input:    "   Hello World   ",
			expected: "Hello World",
		},
		{
			name:     "TrimSpace_StringWithOnlySpaces",
			input:    "   ",
			expected: "",
		},
		{
			name:     "TrimSpace_StringWithNoSpaces",
			input:    "HelloWorld",
			expected: "HelloWorld",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			summary := HtmlSummary{}
			result := summary.trimSpace(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, but got: %s", tc.expected, result)
			}
		})
	}
}
