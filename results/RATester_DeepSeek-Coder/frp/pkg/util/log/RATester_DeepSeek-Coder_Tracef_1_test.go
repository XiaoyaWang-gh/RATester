package log

import (
	"fmt"
	"testing"
)

func TestTracef_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		v      []interface{}
	}{
		{
			name:   "Test Tracef with one argument",
			format: "This is a trace message: %s",
			v:      []interface{}{"test"},
		},
		{
			name:   "Test Tracef with multiple arguments",
			format: "This is a trace message: %s %d",
			v:      []interface{}{"test", 1},
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
			// Add assertions here to verify the output or the state of your system
		})
	}
}
