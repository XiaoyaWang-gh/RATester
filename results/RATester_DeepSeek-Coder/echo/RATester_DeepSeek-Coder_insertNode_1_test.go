package echo

import (
	"fmt"
	"testing"
)

func TestInsertNode_1(t *testing.T) {
	type args struct {
		method string
		path   string
		t      kind
		rm     routeMethod
	}
	tests := []struct {
		name string
		args args
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

			r := &Router{}
			r.insertNode(tt.args.method, tt.args.path, tt.args.t, tt.args.rm)
		})
	}
}
