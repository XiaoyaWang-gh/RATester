package path

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestJoin_3(t *testing.T) {
	ns := &Namespace{
		deps: &deps.Deps{},
	}

	type testCase struct {
		name     string
		elements []any
		expected string
		err      error
	}

	testCases := []testCase{
		{
			name:     "simple join",
			elements: []any{"path", "to", "file"},
			expected: "path/to/file",
			err:      nil,
		},
		{
			name:     "join with slashes",
			elements: []any{"path", "/to", "file"},
			expected: "path/to/file",
			err:      nil,
		},
		{
			name:     "join with empty elements",
			elements: []any{"path", "", "to", "", "file"},
			expected: "path/to/file",
			err:      nil,
		},
		{
			name:     "join with nil elements",
			elements: []any{"path", nil, "to", nil, "file"},
			expected: "path/to/file",
			err:      nil,
		},
		{
			name:     "join with mixed types",
			elements: []any{"path", "to", 123, "file"},
			expected: "path/to/123/file",
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Join(tc.elements...)
			if err != tc.err {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}
			if result != tc.expected {
				t.Errorf("Expected result %s, but got %s", tc.expected, result)
			}
		})
	}
}
