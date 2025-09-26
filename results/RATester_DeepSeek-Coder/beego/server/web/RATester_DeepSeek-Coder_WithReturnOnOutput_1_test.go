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
			name: "TestWithReturnOnOutput_True",
			args: args{ret: true},
			want: WithReturnOnOutput(true),
		},
		{
			name: "TestWithReturnOnOutput_False",
			args: args{ret: false},
			want: WithReturnOnOutput(false),
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
