package log

import (
	"fmt"
	"testing"
)

func TestDebugf_1(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				format: "Test format %s",
				v:      []any{"value"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				format: "Test format %d",
				v:      []any{123},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Debugf(tt.args.format, tt.args.v...)
		})
	}
}
