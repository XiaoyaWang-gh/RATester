package log

import (
	"fmt"
	"testing"
)

func TestDebugf_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		v      []interface{}
	}{
		{
			name:   "Test with string",
			format: "This is a test %s",
			v:      []interface{}{"string"},
		},
		{
			name:   "Test with int",
			format: "This is a test %d",
			v:      []interface{}{123},
		},
		{
			name:   "Test with float",
			format: "This is a test %f",
			v:      []interface{}{123.456},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Debugf(tt.format, tt.v...)
		})
	}
}
