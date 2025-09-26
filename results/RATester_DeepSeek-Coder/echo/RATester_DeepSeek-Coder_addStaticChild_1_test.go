package echo

import (
	"fmt"
	"testing"
)

func TestAddStaticChild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{
		staticChildren: make(children, 0),
	}

	c := &node{
		prefix: "/test",
	}

	n.addStaticChild(c)

	if len(n.staticChildren) != 1 {
		t.Errorf("Expected 1 child, got %d", len(n.staticChildren))
	}

	if n.staticChildren[0].prefix != "/test" {
		t.Errorf("Expected child prefix to be '/test', got '%s'", n.staticChildren[0].prefix)
	}
}
