package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddtree_1(t *testing.T) {
	type args struct {
		segments  []string
		tree      *Tree
		wildcards []string
		reg       string
	}
	tests := []struct {
		name string
		args args
		want *Tree
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

			tree := &Tree{}
			tree.addtree(tt.args.segments, tt.args.tree, tt.args.wildcards, tt.args.reg)
			if !reflect.DeepEqual(tree, tt.want) {
				t.Errorf("addtree() = %v, want %v", tree, tt.want)
			}
		})
	}
}
