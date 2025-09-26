package binder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEqualFieldType_1(t *testing.T) {
	type args struct {
		out  any
		kind reflect.Kind
		key  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test equalFieldType",
			args: args{
				out:  nil,
				kind: reflect.Kind(0),
				key:  "",
			},
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

			if got := equalFieldType(tt.args.out, tt.args.kind, tt.args.key); got != tt.want {
				t.Errorf("equalFieldType() = %v, want %v", got, tt.want)
			}
		})
	}
}
