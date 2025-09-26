package parse

import (
	"fmt"
	"testing"
)

func TestparseDefinition_1(t *testing.T) {
	type args struct {
		tree *Tree
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

			tt.args.tree.parseDefinition()
		})
	}
}
