package tpl

import (
	"fmt"
	"testing"
)

func TestExtractBaseOf_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "error: failed to open file: no such file or directory",
			expected: "no such file or directory",
		},
		{
			name:     "Test case 2",
			input:    "error: failed to open file: permission denied",
			expected: "permission denied",
		},
		{
			name:     "Test case 3",
			input:    "error: failed to open file: file is a directory",
			expected: "file is a directory",
		},
		{
			name:     "Test case 4",
			input:    "error: failed to open file: too many open files",
			expected: "too many open files",
		},
		{
			name:     "Test case 5",
			input:    "error: failed to open file: invalid argument",
			expected: "invalid argument",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := extractBaseOf(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
