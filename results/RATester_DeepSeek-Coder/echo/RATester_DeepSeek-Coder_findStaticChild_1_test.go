package echo

import (
	"fmt"
	"testing"
)

func TestFindStaticChild_1(t *testing.T) {
	type test struct {
		name     string
		children []*node
		label    byte
		want     *node
	}

	tests := []test{
		{
			name: "Test case 1",
			children: []*node{
				{label: 'a', isLeaf: true},
				{label: 'b', isLeaf: true},
				{label: 'c', isLeaf: true},
			},
			label: 'b',
			want:  &node{label: 'b', isLeaf: true},
		},
		{
			name: "Test case 2",
			children: []*node{
				{label: 'a', isLeaf: true},
				{label: 'b', isLeaf: true},
				{label: 'c', isLeaf: true},
			},
			label: 'd',
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			n := &node{
				staticChildren: tt.children,
			}

			for _, c := range n.staticChildren {
				c.addStaticChild(c)
			}

			got := n.findStaticChild(tt.label)

			if got != tt.want {
				t.Errorf("findStaticChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
