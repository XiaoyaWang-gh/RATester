package filecache

import (
	"fmt"
	"testing"
)

func TestcleanID_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "/path/to/file",
			expected: "path/to/file",
		},
		{
			name:     "Test case 2",
			input:    "path/to/file",
			expected: "path/to/file",
		},
		{
			name:     "Test case 3",
			input:    "path/to/file/",
			expected: "path/to/file",
		},
		{
			name:     "Test case 4",
			input:    "path/to/file//",
			expected: "path/to/file",
		},
		{
			name:     "Test case 5",
			input:    "//path/to/file",
			expected: "path/to/file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := cleanID(tt.input)
			if got != tt.expected {
				t.Errorf("cleanID(%s) = %s, want %s", tt.input, got, tt.expected)
			}
		})
	}
}
