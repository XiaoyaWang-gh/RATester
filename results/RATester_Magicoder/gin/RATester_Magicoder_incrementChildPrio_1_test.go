package gin

import (
	"fmt"
	"testing"
)

func TestincrementChildPrio_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{
		path:    "/",
		indices: "abc",
		children: []*node{
			{path: "a", priority: 1},
			{path: "b", priority: 2},
			{path: "c", priority: 3},
		},
	}

	newPos := n.incrementChildPrio(1)

	if newPos != 1 {
		t.Errorf("Expected newPos to be 1, but got %d", newPos)
	}

	if n.children[1].priority != 3 {
		t.Errorf("Expected priority of second child to be 3, but got %d", n.children[1].priority)
	}

	if n.indices != "abc" {
		t.Errorf("Expected indices to be 'abc', but got '%s'", n.indices)
	}
}
