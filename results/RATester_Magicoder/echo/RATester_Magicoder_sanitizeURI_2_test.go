package echo

import (
	"fmt"
	"testing"
)

func TestSanitizeURI_2(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single slash", "/", "/"},
		{"double slash", "//", "/"},
		{"leading slash", "/test", "/test"},
		{"double leading slash", "//test", "/test"},
		{"leading backslash", "\\test", "/test"},
		{"double leading backslash", "\\\\test", "/test"},
		{"leading backslash and slash", "\\/test", "/test"},
		{"double leading backslash and slash", "\\\\/test", "/test"},
		{"no leading slash or backslash", "test", "test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := sanitizeURI(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, result)
			}
		})
	}
}
