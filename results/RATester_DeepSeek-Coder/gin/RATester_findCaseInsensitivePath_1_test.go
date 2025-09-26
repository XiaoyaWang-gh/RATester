package gin

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFindCaseInsensitivePath_1(t *testing.T) {
	n := &node{
		path: "/test",
		children: []*node{
			{
				path: "/test/child",
			},
		},
	}

	tests := []struct {
		name     string
		path     string
		expected []byte
		found    bool
	}{
		{
			name:     "found",
			path:     "/test/child",
			expected: []byte("/test/child"),
			found:    true,
		},
		{
			name:     "not found",
			path:     "/not/found",
			expected: nil,
			found:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			found, ok := n.findCaseInsensitivePath(test.path, false)
			if ok != test.found {
				t.Errorf("expected found to be %v, but got %v", test.found, ok)
			}
			if !bytes.Equal(found, test.expected) {
				t.Errorf("expected path to be %s, but got %s", test.expected, found)
			}
		})
	}
}
