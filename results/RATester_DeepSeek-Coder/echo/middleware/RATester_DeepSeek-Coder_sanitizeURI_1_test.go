package middleware

import (
	"fmt"
	"testing"
)

func TestSanitizeURI_1(t *testing.T) {

	testCases := []struct {
		name     string
		uri      string
		expected string
	}{
		{"Empty URI", "", "/"},
		{"Single slash URI", "/", "/"},
		{"Double slash URI", "//", "/"},
		{"URI with single slash", "//test", "/test"},
		{"URI with double slash", "//test/test", "/test/test"},
		{"URI with single backslash", "\\test", "/test"},
		{"URI with double backslash", "\\\\test", "/test"},
		{"URI with single backslash and double slash", "\\/test", "/test"},
		{"URI with double backslash and double slash", "\\\\test", "/test"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := sanitizeURI(tc.uri)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
