package framework

import (
	"fmt"
	"testing"
)

func TestLogf_2(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{
			name:   "Test with no arguments",
			format: "This is a test",
			args:   []interface{}{},
		},
		{
			name:   "Test with one argument",
			format: "This is a test with argument: %v",
			args:   []interface{}{"test"},
		},
		{
			name:   "Test with multiple arguments",
			format: "This is a test with arguments: %v, %v",
			args:   []interface{}{"test1", "test2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Logf(tt.format, tt.args...)
		})
	}
}
