package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestindirectInterface_3(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
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

			if got := indirectInterface(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirectInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
