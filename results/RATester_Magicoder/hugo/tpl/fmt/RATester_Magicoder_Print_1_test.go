package fmt

import (
	"fmt"
	"testing"
)

func TestPrint_1(t *testing.T) {
	ns := &Namespace{}

	testCases := []struct {
		name     string
		args     []any
		expected string
	}{
		{
			name:     "single string",
			args:     []any{"hello"},
			expected: "hello",
		},
		{
			name:     "multiple strings",
			args:     []any{"hello", "world"},
			expected: "helloworld",
		},
		{
			name:     "mixed types",
			args:     []any{"hello", 123, true},
			expected: "hello123true",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ns.Print(tc.args...)
			if result != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
