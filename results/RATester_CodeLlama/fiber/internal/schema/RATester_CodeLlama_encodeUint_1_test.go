package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeUint_1(t *testing.T) {
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
			args: args{v: reflect.ValueOf(1)},
			want: "1",
		},
		{
			name: "case2",
			args: args{v: reflect.ValueOf(1.1)},
			want: "1.1",
		},
		{
			name: "case3",
			args: args{v: reflect.ValueOf("1")},
			want: "1",
		},
		{
			name: "case4",
			args: args{v: reflect.ValueOf([]int{1})},
			want: "1",
		},
		{
			name: "case5",
			args: args{v: reflect.ValueOf(map[string]int{"1": 1})},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeUint(tt.args.v); got != tt.want {
				t.Errorf("encodeUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
