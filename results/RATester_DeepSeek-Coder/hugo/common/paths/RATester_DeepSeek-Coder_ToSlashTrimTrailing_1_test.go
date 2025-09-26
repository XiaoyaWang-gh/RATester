package paths

import (
	"fmt"
	"testing"
)

func TestToSlashTrimTrailing_1(t *testing.T) {
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
			name:     "no trailing slash",
			input:    "path/to/dir",
			expected: "path/to/dir",
		},
		{
			name:     "trailing slash",
			input:    "path/to/dir/",
			expected: "path/to/dir",
		},
		{
			name:     "multiple trailing slashes",
			input:    "path/to/dir////////",
			expected: "path/to/dir",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ToSlashTrimTrailing(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
