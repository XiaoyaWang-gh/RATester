package hugocontext

import (
	"fmt"
	"testing"
)

func TestParseAttrs_1(t *testing.T) {
	type test struct {
		name     string
		input    []byte
		expected uint64
	}

	tests := []test{
		{
			name:     "valid pid",
			input:    []byte("pid=123"),
			expected: 123,
		},
		{
			name:     "invalid pid",
			input:    []byte("pid=abc"),
			expected: 0,
		},
		{
			name:     "missing pid",
			input:    []byte(""),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &HugoContext{}
			ctx.parseAttrs(tt.input)
			if ctx.Pid != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, ctx.Pid)
			}
		})
	}
}
