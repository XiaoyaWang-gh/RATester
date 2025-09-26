package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeString_1(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test encodeString",
			args: args{
				v: reflect.ValueOf("test"),
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeString(tt.args.v); got != tt.want {
				t.Errorf("encodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
