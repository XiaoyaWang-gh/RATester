package gin

import (
	"fmt"
	"testing"
)

func TestaddChild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{
		path:     "/",
		children: []*node{},
	}

	child := &node{
		path: "/child",
	}

	n.addChild(child)

	if len(n.children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(n.children))
	}

	if n.children[0] != child {
		t.Errorf("Expected child to be %v, got %v", child, n.children[0])
	}
}
