package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestencodeFloat_1(t *testing.T) {
	type args struct {
		v    reflect.Value
		bits int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
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
