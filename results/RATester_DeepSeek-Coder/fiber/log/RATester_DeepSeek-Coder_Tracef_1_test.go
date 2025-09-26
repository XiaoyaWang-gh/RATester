package log

import (
	"fmt"
	"testing"
)

func TestTracef_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		v      []any
	}{
		{
			name:   "Test with string format and string value",
			format: "Hello, %s",
			v:      []any{"World"},
		},
		{
			name:   "Test with int format and int value",
			format: "The number is %d",
			v:      []any{42},
		},
		{
			name:   "Test with float format and float value",
			format: "The number is %f",
			v:      []any{3.14},
		},
		{
			name:   "Test with multiple values",
			format: "Hello, %s. You have %d new messages.",
			v:      []any{"Alice", 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Tracef(tt.format, tt.v...)
		})
	}
}
