package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeFloat_1(t *testing.T) {
	type args struct {
		v    reflect.Value
		bits int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{
				v:    reflect.ValueOf(1.1),
				bits: 64,
			},
			want: "1.100000",
		},
		{
			name: "case 2",
			args: args{
				v:    reflect.ValueOf(1.1),
				bits: 32,
			},
			want: "1.100000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeFloat(tt.args.v, tt.args.bits); got != tt.want {
				t.Errorf("encodeFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
