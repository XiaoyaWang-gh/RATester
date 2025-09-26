package paths

import (
	"fmt"
	"testing"
)

func TestBaseNameNoIdentifier_1(t *testing.T) {
	type test struct {
		name     string
		path     string
		expected string
	}

	tests := []test{
		{
			name:     "Bundle",
			path:     "/path/to/bundle/",
			expected: "bundle",
		},
		{
			name:     "Content",
			path:     "/path/to/content.md",
			expected: "content",
		},
		{
			name:     "Empty",
			path:     "",
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Path{s: tc.path}
			result := p.BaseNameNoIdentifier()
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
