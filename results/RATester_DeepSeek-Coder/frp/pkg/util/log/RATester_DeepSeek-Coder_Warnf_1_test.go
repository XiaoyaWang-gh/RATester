package log

import (
	"fmt"
	"testing"
)

func TestWarnf_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		v      []interface{}
	}{
		{
			name:   "TestWarnf_0",
			format: "This is a warning: %s",
			v:      []interface{}{"TestWarnf_0"},
		},
		{
			name:   "TestWarnf_1",
			format: "This is a warning: %d",
			v:      []interface{}{123456},
		},
		{
			name:   "TestWarnf_2",
			format: "This is a warning: %v",
			v:      []interface{}{"TestWarnf_2"},
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
		})
	}
}
