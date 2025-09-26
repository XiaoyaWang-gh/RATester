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
			name:   "Test case 1",
			format: "This is a test %s",
			args:   []interface{}{"message"},
		},
		{
			name:   "Test case 2",
			format: "This is another test %d",
			args:   []interface{}{123},
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
