package log

import (
	"fmt"
	"testing"
)

func TestFatalf_1(t *testing.T) {
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
				format: "test %s",
				v:      []any{"value"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				format: "test %d",
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

			Fatalf(tt.args.format, tt.args.v...)
		})
	}
}
