package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestencodeInt_1(t *testing.T) {
	type args struct {
		v reflect.Value
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

			if got := encodeInt(tt.args.v); got != tt.want {
				t.Errorf("encodeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
