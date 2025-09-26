package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestgetRootsReverse_1(t *testing.T) {
	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
		realMapToRoot: radix.New(),
	}

	testCases := []struct {
		name     string
		key      string
		expected string
		roots    []RootMapping
	}{
		{
			name:     "Test case 1",
			key:      "test1",
			expected: "test1",
			roots: []RootMapping{
				{
					From: "test1",
				},
			},
		},
		{
			name:     "Test case 2",
			key:      "test2",
			expected: "test2",
			roots: []RootMapping{
				{
					From: "test2",
				},
			},
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
				fs.realMapToRoot.Insert(root.From, root)
			}

			actual, _ := fs.getRootsReverse(tc.key)
			if actual != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, actual)
			}
		})
	}
}
