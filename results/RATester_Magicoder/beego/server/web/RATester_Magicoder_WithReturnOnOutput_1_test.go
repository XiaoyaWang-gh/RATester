package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithReturnOnOutput_1(t *testing.T) {
	type args struct {
		ret bool
	}
	tests := []struct {
		name string
		args args
		want FilterOpt
	}{
		{
			name: "Test case 1",
			args: args{
				ret: true,
			},
			want: func(opts *filterOpts) {
				opts.returnOnOutput = true
			},
		},
		{
			name: "Test case 2",
			args: args{
				ret: false,
			},
			want: func(opts *filterOpts) {
				opts.returnOnOutput = false
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

			if got := WithReturnOnOutput(tt.args.ret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithReturnOnOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
