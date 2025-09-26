package paths

import (
	"fmt"
	"testing"
)

func TestNameNoLang_1(t *testing.T) {
	type test struct {
		name     string
		path     string
		expected string
	}

	tests := []test{
		{
			name:     "No identifier",
			path:     "/path/to/file.txt",
			expected: "file.txt",
		},
		{
			name:     "With identifier",
			path:     "/path/to/file.txt@identifier",
			expected: "file.txt",
		},
		{
			name:     "With identifier and lang",
			path:     "/path/to/file.txt@identifier@lang",
			expected: "file.txt",
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

			got := p.NameNoLang()
			if got != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, got)
			}
		})
	}
}
