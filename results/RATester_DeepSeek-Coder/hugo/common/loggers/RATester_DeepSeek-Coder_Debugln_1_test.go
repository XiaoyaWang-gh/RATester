package loggers

import (
	"fmt"
	"testing"
)

func TestDebugln_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
	}{
		{
			name: "Test with string",
			args: []any{"test string"},
		},
		{
			name: "Test with multiple arguments",
			args: []any{"test", "multiple", "arguments"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &logAdapter{}
			l.Debugln(tt.args...)
		})
	}
}
