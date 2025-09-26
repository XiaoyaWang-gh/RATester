package path

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestJoin_3(t *testing.T) {
	t.Parallel()

	ns := &Namespace{
		deps: &deps.Deps{},
	}

	type testCase struct {
		name     string
		elements []any
		expected string
		wantErr  bool
	}

	tests := []testCase{
		{
			name:     "Join with string elements",
			elements: []any{"path", "to", "file.txt"},
			expected: "path/to/file.txt",
			wantErr:  false,
		},
		{
			name:     "Join with string slice elements",
			elements: []any{[]string{"path", "to", "file.txt"}},
			expected: "path/to/file.txt",
			wantErr:  false,
		},
		{
			name:     "Join with mixed elements",
			elements: []any{"path", []string{"to", "file.txt"}},
			expected: "path/to/file.txt",
			wantErr:  false,
		},
		{
			name:     "Join with invalid elements",
			elements: []any{123},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Join(tt.elements...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Join() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("Namespace.Join() = %v, want %v", got, tt.expected)
			}
		})
	}
}
