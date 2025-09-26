package log

import (
	"fmt"
	"testing"
)

func TestInfof_1(t *testing.T) {
	type args struct {
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestInfof_0",
			args: args{
				format: "TestInfof_0",
				v:      []any{"TestInfof_0"},
			},
		},
		{
			name: "TestInfof_1",
			args: args{
				format: "TestInfof_1",
				v:      []any{"TestInfof_1"},
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

			Infof(tt.args.format, tt.args.v...)
		})
	}
}
