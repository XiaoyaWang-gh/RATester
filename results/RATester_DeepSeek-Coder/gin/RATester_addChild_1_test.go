package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddChild_1(t *testing.T) {
	type args struct {
		child *node
	}
	tests := []struct {
		name string
		n    *node
		args args
		want []*node
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

			n := &node{
				path:      tt.n.path,
				indices:   tt.n.indices,
				wildChild: tt.n.wildChild,
				children:  tt.n.children,
			}
			n.addChild(tt.args.child)
			if !reflect.DeepEqual(n.children, tt.want) {
				t.Errorf("node.addChild() = %v, want %v", n.children, tt.want)
			}
		})
	}
}
