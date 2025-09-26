package log

import (
	"fmt"
	"testing"
)

func TestTracef_1(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Tracef with string and int",
			args: args{
				format: "Test %s with %d",
				v:      []any{"Tracef", 1},
			},
		},
		{
			name: "Test Tracef with string and float",
			args: args{
				format: "Test %s with %f",
				v:      []any{"Tracef", 1.23},
			},
		},
		{
			name: "Test Tracef with string and bool",
			args: args{
				format: "Test %s with %t",
				v:      []any{"Tracef", true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Tracef(tt.args.format, tt.args.v...)
		})
	}
}
