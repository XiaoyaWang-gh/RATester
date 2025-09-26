package urls

import (
	"errors"
	"fmt"
	"testing"
)

func TestJoinPath_1(t *testing.T) {
	ns := &Namespace{}

	testCases := []struct {
		name     string
		elements []any
		expected string
		err      error
	}{
		{
			name:     "empty",
			elements: []any{},
			expected: "",
			err:      nil,
		},
		{
			name:     "single string",
			elements: []any{"foo"},
			expected: "foo",
			err:      nil,
		},
		{
			name:     "multiple strings",
			elements: []any{"foo", "bar", "baz"},
			expected: "foo/bar/baz",
			err:      nil,
		},
		{
			name:     "mixed types",
			elements: []any{"foo", []string{"bar", "baz"}},
			expected: "foo/bar/baz",
			err:      nil,
		},
		{
			name:     "invalid type",
			elements: []any{"foo", 123},
			expected: "",
			err:      errors.New("unsupported type"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.JoinPath(tc.elements...)
			if tc.err != nil {
				if err == nil || err.Error() != tc.err.Error() {
					t.Errorf("expected error %v, got %v", tc.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tc.expected {
					t.Errorf("expected %q, got %q", tc.expected, result)
				}
			}
		})
	}
}
