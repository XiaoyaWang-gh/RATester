package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsValidStructPointer_1(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				v: reflect.ValueOf(struct{}{}),
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				v: reflect.ValueOf(struct{}{}),
			},
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

			if got := isValidStructPointer(tt.args.v); got != tt.want {
				t.Errorf("isValidStructPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
