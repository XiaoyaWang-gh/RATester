package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeBool_1(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{v: reflect.ValueOf(true)},
			want: "true",
		},
		{
			name: "case2",
			args: args{v: reflect.ValueOf(false)},
			want: "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeBool(tt.args.v); got != tt.want {
				t.Errorf("encodeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
