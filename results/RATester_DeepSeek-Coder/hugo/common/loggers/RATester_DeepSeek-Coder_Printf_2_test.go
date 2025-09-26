package loggers

import (
	"fmt"
	"strings"
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
			name:     "TestPrintf_0",
			format:   "test %s",
			args:     []any{"test"},
			expected: "test test\n",
		},
		{
			name:     "TestPrintf_1",
			format:   "test %d",
			args:     []any{123},
			expected: "test 123\n",
		},
		{
			name:     "TestPrintf_2",
			format:   "test %v",
			args:     []any{struct{}{}},
			expected: "test {}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			out := &strings.Builder{}
			l := &logAdapter{out: out}

			l.Printf(tt.format, tt.args...)

			if got := out.String(); got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
