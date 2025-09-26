package log

import (
	"fmt"
	"testing"
)

func TestTrace_2(t *testing.T) {
	tests := []struct {
		name string
		args []any
	}{
		{
			name: "Test Trace with string",
			args: []any{"test"},
		},
		{
			name: "Test Trace with int",
			args: []any{1},
		},
		{
			name: "Test Trace with float",
			args: []any{1.1},
		},
		{
			name: "Test Trace with bool",
			args: []any{true},
		},
		{
			name: "Test Trace with nil",
			args: []any{nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Trace(tt.args...)
		})
	}
}
