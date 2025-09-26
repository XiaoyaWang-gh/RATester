package fmt

import (
	"fmt"
	"testing"
)

func TestPrintf_1(t *testing.T) {
	ns := &Namespace{}

	testCases := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{
			name:     "Simple string",
			format:   "Hello, %s",
			args:     []any{"World"},
			expected: "Hello, World",
		},
		{
			name:     "Multiple arguments",
			format:   "%s %s",
			args:     []any{"Hello", "World"},
			expected: "Hello World",
		},
		{
			name:     "Integer",
			format:   "%d",
			args:     []any{42},
			expected: "42",
		},
		{
			name:     "Float",
			format:   "%f",
			args:     []any{3.14},
			expected: "3.140000",
		},
		{
			name:     "Boolean",
			format:   "%t",
			args:     []any{true},
			expected: "true",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ns.Printf(tc.format, tc.args...)
			if result != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
