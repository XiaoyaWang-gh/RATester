package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddMethod_1(t *testing.T) {
	type args struct {
		method string
		h      *routeMethod
	}
	tests := []struct {
		name string
		n    *node
		args args
		want *node
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

			n := &node{}
			n.addMethod(tt.args.method, tt.args.h)
			if !reflect.DeepEqual(n, tt.want) {
				t.Errorf("node.addMethod() = %v, want %v", n, tt.want)
			}
		})
	}
}
