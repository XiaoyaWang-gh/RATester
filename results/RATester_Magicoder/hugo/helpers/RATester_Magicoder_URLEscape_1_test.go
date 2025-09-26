package helpers

import (
	"fmt"
	"testing"
)

func TestURLEscape_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "URL with special characters",
			input:    "http://example.com/path?query=value#fragment",
			expected: "http://example.com/path?query=value#fragment",
		},
		{
			name:     "URL with no special characters",
			input:    "http://example.com/path",
			expected: "http://example.com/path",
		},
		{
			name:     "URL with space",
			input:    "http://example.com/path with space",
			expected: "http://example.com/path%20with%20space",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ps := &PathSpec{}
			result := ps.URLEscape(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
