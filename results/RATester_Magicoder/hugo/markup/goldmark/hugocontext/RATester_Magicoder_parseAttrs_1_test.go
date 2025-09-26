package hugocontext

import (
	"fmt"
	"testing"
)

func TestparseAttrs_1(t *testing.T) {
	tests := []struct {
		name     string
		attrs    string
		expected uint64
	}{
		{
			name:     "valid pid",
			attrs:    "pid=123",
			expected: 123,
		},
		{
			name:     "invalid pid",
			attrs:    "pid=abc",
			expected: 0,
		},
		{
			name:     "empty pid",
			attrs:    "pid=",
			expected: 0,
		},
		{
			name:     "no pid",
			attrs:    "",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &HugoContext{}
			ctx.parseAttrs([]byte(test.attrs))
			if ctx.Pid != test.expected {
				t.Errorf("Expected Pid to be %d, but got %d", test.expected, ctx.Pid)
			}
		})
	}
}
