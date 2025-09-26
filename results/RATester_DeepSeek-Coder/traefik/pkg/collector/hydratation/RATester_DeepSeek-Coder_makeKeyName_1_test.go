package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeKeyName_1(t *testing.T) {
	type args struct {
		typ reflect.Type
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

			if got := makeKeyName(tt.args.typ); got != tt.want {
				t.Errorf("makeKeyName() = %v, want %v", got, tt.want)
			}
		})
	}
}
