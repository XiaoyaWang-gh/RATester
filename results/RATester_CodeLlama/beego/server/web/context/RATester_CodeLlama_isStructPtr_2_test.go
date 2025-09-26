package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsStructPtr_2(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{t: reflect.TypeOf(struct{}{})},
			want: false,
		},
		{
			name: "case2",
			args: args{t: reflect.TypeOf(&struct{}{})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isStructPtr(tt.args.t); got != tt.want {
				t.Errorf("isStructPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
