package log

import (
	"fmt"
	"testing"
)

func TestWarnf_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		v      []any
	}{
		{
			name:   "Test with string format and string value",
			format: "This is a warning: %s",
			v:      []any{"warning message"},
		},
		{
			name:   "Test with int format and int value",
			format: "This is a warning: %d",
			v:      []any{123},
		},
		{
			name:   "Test with float format and float value",
			format: "This is a warning: %f",
			v:      []any{123.456},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Warnf(tt.format, tt.v...)
			// Add assertions here to verify the behavior of the function under test.
		})
	}
}
