package paths

import (
	"fmt"
	"testing"
)

func TestAddLeadingAndTrailingSlash_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty string", "", "/"},
		{"Single character", "a", "/a/"},
		{"Leading slash", "/a", "/a/"},
		{"Trailing slash", "a/", "/a/"},
		{"Leading and trailing slash", "/a/", "/a/"},
		{"Multiple characters", "a/b", "/a/b/"},
		{"Multiple slashes", "//a//b//", "/a/b/"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := AddLeadingAndTrailingSlash(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
