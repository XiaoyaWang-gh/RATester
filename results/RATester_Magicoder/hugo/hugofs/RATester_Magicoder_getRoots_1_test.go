package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestgetRoots_1(t *testing.T) {
	fs := &RootMappingFs{
		rootMapToReal: radix.New(),
	}

	testCases := []struct {
		name     string
		key      string
		expected string
		roots    []RootMapping
	}{
		{
			name:     "simple case",
			key:      "/foo/bar",
			expected: "/foo",
			roots: []RootMapping{
				{From: "/foo"},
			},
		},
		{
			name:     "multiple roots",
			key:      "/foo/bar",
			expected: "/foo",
			roots: []RootMapping{
				{From: "/foo"},
				{From: "/bar"},
			},
		},
		{
			name:     "no roots",
			key:      "/foo/bar",
			expected: "",
			roots:    []RootMapping{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			for _, root := range tc.roots {
				fs.rootMapToReal.Insert(root.From, root)
			}

			actual, roots := fs.getRoots(tc.key)

			if actual != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, actual)
			}

			if len(roots) != len(tc.roots) {
				t.Errorf("Expected %d roots, but got %d", len(tc.roots), len(roots))
			}

			for i, root := range roots {
				if root.From != tc.roots[i].From {
					t.Errorf("Expected root %s, but got %s", tc.roots[i].From, root.From)
				}
			}
		})
	}
}
