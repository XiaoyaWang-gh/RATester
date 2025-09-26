package glob

import (
	"fmt"
	"testing"
)

func TestNormalizeFilenameGlobPattern_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with absolute path",
			input:    "/test/path/*",
			expected: "/test/path/*",
		},
		{
			name:     "Test with relative path",
			input:    "test/path/*",
			expected: "/test/path/*",
		},
		{
			name:     "Test with no leading slash",
			input:    "test/path/file",
			expected: "/test/path/file",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := normalizeFilenameGlobPattern(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
