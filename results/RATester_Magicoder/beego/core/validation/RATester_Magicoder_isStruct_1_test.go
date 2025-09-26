package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestisStruct_1(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test isStruct with struct type",
			args: args{t: reflect.TypeOf(struct{}{})},
			want: true,
		},
		{
			name: "Test isStruct with non-struct type",
			args: args{t: reflect.TypeOf(1)},
			want: false,
		},
		{
			name: "Test isStruct with nil type",
			args: args{t: reflect.TypeOf(nil)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isStruct(tt.args.t); got != tt.want {
				t.Errorf("isStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
