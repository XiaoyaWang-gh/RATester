package gin

import (
	"fmt"
	"testing"
)

func TestInsertChild_1(t *testing.T) {
	type args struct {
		path     string
		fullPath string
		handlers HandlersChain
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

			n := &node{}
			n.insertChild(tt.args.path, tt.args.fullPath, tt.args.handlers)
		})
	}
}
