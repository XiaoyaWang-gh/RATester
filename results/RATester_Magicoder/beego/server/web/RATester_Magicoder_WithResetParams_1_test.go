package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithResetParams_1(t *testing.T) {
	type args struct {
		reset bool
	}
	tests := []struct {
		name string
		args args
		want FilterOpt
	}{
		{
			name: "Test case 1",
			args: args{
				reset: true,
			},
			want: func(opts *filterOpts) {
				opts.resetParams = true
			},
		},
		{
			name: "Test case 2",
			args: args{
				reset: false,
			},
			want: func(opts *filterOpts) {
				opts.resetParams = false
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

			if got := WithResetParams(tt.args.reset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithResetParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
