package gin

import (
	"fmt"
	"testing"
)

func TestIncrementChildPrio_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{
		path:    "/test",
		indices: "012",
		children: []*node{
			{priority: 1},
			{priority: 2},
			{priority: 3},
		},
	}

	newPos := n.incrementChildPrio(1)

	if newPos != 1 {
		t.Errorf("Expected new position to be 1, got %d", newPos)
	}

	if n.children[1].priority != 3 {
		t.Errorf("Expected priority of second child to be 3, got %d", n.children[1].priority)
	}

	if n.indices != "012" {
		t.Errorf("Expected indices to be '012', got '%s'", n.indices)
	}
}
