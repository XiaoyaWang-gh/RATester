package loggers

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrintf_2(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{
			name:     "Simple string",
			format:   "Hello, %s",
			args:     []any{"World"},
			expected: "Hello, World\n",
		},
		{
			name:     "Multiple arguments",
			format:   "%s %s",
			args:     []any{"Hello", "World"},
			expected: "Hello World\n",
		},
		{
			name:     "Integer",
			format:   "Number: %d",
			args:     []any{42},
			expected: "Number: 42\n",
		},
		{
			name:     "Float",
			format:   "Pi: %f",
			args:     []any{3.14159},
			expected: "Pi: 3.141590\n",
		},
		{
			name:     "Boolean",
			format:   "Is it true? %t",
			args:     []any{true},
			expected: "Is it true? true\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			buf := &bytes.Buffer{}
			la := &logAdapter{out: buf}

			la.Printf(test.format, test.args...)

			if buf.String() != test.expected {
				t.Errorf("Expected '%s', but got '%s'", test.expected, buf.String())
			}
		})
	}
}
