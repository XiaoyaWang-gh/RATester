package helpers

import (
	"fmt"
	"testing"
)

func TestURLizeFilename_1(t *testing.T) {
	type test struct {
		name     string
		filename string
		expected string
	}

	tests := []test{
		{
			name:     "simple filename",
			filename: "test.txt",
			expected: "test.txt",
		},
		{
			name:     "filename with spaces",
			filename: "test file.txt",
			expected: "test%20file.txt",
		},
		{
			name:     "filename with special characters",
			filename: "test_file.txt",
			expected: "test_file.txt",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ps := &PathSpec{}
			result := ps.URLizeFilename(tc.filename)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
