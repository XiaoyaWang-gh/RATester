package stripprefixregex

import (
	"fmt"
	"testing"
)

func TestEnsureLeadingSlash_2(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", ""},
		{"String with leading slash", "/test", "/test"},
		{"String without leading slash", "test", "/test"},
		{"String with leading space", " test", "/ test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ensureLeadingSlash(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
