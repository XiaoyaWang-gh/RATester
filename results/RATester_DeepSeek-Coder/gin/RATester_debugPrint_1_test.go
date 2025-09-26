package gin

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDebugPrint_1(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		values   []any
		expected string
	}{
		{
			name:     "Test with no values",
			format:   "This is a test",
			values:   []any{},
			expected: "[GIN-debug] This is a test\n",
		},
		{
			name:     "Test with one value",
			format:   "This is a test: %v",
			values:   []any{123},
			expected: "[GIN-debug] This is a test: 123\n",
		},
		{
			name:     "Test with multiple values",
			format:   "This is a test: %v, %v, %v",
			values:   []any{123, "abc", true},
			expected: "[GIN-debug] This is a test: 123, abc, true\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &bytes.Buffer{}
			DefaultWriter = w
			debugPrint(tt.format, tt.values...)
			got := w.String()
			if got != tt.expected {
				t.Errorf("Expected '%s', but got '%s'", tt.expected, got)
			}
		})
	}
}
