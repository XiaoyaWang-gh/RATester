package deployconfig

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatches_1(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		path     string
		expected bool
	}{
		{
			name:     "Matching path",
			pattern:  ".*",
			path:     "/test/path",
			expected: true,
		},
		{
			name:     "Non-matching path",
			pattern:  "^/test/path$",
			path:     "/not/matching",
			expected: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Matcher{
				Pattern: tt.pattern,
				Re:      regexp.MustCompile(tt.pattern),
			}

			got := m.Matches(tt.path)
			if got != tt.expected {
				t.Errorf("Matches() = %v, want %v", got, tt.expected)
			}
		})
	}
}
