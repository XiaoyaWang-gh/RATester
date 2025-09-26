package paths

import (
	"fmt"
	"testing"
)

func TestToSlashTrim_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"root path", "/", ""},
		{"sub path", "/sub/path", "sub/path"},
		{"sub path with trailing slash", "/sub/path/", "sub/path"},
		{"sub path with leading slash", "/sub/path/", "sub/path"},
		{"sub path with leading and trailing slash", "//sub/path//", "sub/path"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ToSlashTrim(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
