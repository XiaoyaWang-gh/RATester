package paths

import (
	"fmt"
	"testing"
)

func TestNameNoExt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "no identifier",
			path:     "/path/to/file.txt",
			expected: "file",
		},
		{
			name:     "with identifier",
			path:     "/path/to/file-identifier.txt",
			expected: "file-identifier",
		},
		{
			name:     "no extension",
			path:     "/path/to/file",
			expected: "file",
		},
		{
			name:     "empty path",
			path:     "",
			expected: "",
		},
	}

	for _, test := range tests {
		p := &Path{s: test.path}
		got := p.NameNoExt()
		if got != test.expected {
			t.Errorf("Test %q: expected %q, got %q", test.name, test.expected, got)
		}
	}
}
