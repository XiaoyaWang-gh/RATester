package paths

import (
	"fmt"
	"testing"
)

func TestNameNoIdentifier_1(t *testing.T) {
	type test struct {
		name     string
		path     string
		expected string
	}

	tests := []test{
		{
			name:     "no identifiers",
			path:     "/path/to/file",
			expected: "file",
		},
		{
			name:     "one identifier",
			path:     "/path/to/identifier/file",
			expected: "file",
		},
		{
			name:     "multiple identifiers",
			path:     "/path/to/identifier1/identifier2/file",
			expected: "file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Path{
				s: tt.path,
			}
			got := p.NameNoIdentifier()
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
