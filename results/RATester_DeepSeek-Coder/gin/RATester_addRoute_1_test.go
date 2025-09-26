package gin

import (
	"fmt"
	"testing"
)

func TestAddRoute_1(t *testing.T) {
	type args struct {
		path     string
		handlers HandlersChain
	}
	tests := []struct {
		name string
		n    *node
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

			tt.n.addRoute(tt.args.path, tt.args.handlers)
		})
	}
}
