package deployconfig

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatches_1(t *testing.T) {
	type testCase struct {
		name     string
		pattern  string
		path     string
		expected bool
	}

	testCases := []testCase{
		{
			name:     "Matching path",
			pattern:  "^/path/to/file$",
			path:     "/path/to/file",
			expected: true,
		},
		{
			name:     "Non-matching path",
			pattern:  "^/path/to/file$",
			path:     "/path/to/another/file",
			expected: false,
		},
		{
			name:     "Matching path with regex special characters",
			pattern:  "^/path/to/[a-z]+/file$",
			path:     "/path/to/some/file",
			expected: true,
		},
		{
			name:     "Non-matching path with regex special characters",
			pattern:  "^/path/to/[a-z]+/file$",
			path:     "/path/to/SOME/file",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			re, err := regexp.Compile(tc.pattern)
			if err != nil {
				t.Fatalf("Failed to compile regexp: %v", err)
			}

			m := &Matcher{
				Pattern: tc.pattern,
				Re:      re,
			}

			result := m.Matches(tc.path)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
