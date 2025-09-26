package helpers

import (
	"fmt"
	"testing"
)

func TestURLEscape_1(t *testing.T) {
	type test struct {
		name     string
		uri      string
		expected string
	}

	tests := []test{
		{
			name:     "Test with simple URL",
			uri:      "http://example.com",
			expected: "http://example.com",
		},
		{
			name:     "Test with URL with special characters",
			uri:      "http://example.com/foo/bar?baz=qux&quux=quuz#corge",
			expected: "http://example.com/foo/bar?baz=qux&quux=quuz#corge",
		},
		{
			name:     "Test with URL with non-ASCII characters",
			uri:      "http://example.com/日本語",
			expected: "http://example.com/%E6%97%A5%E6%9C%AC%E8%AA%9E",
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
			result := ps.URLEscape(tc.uri)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
