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

	n := &node{}
	c := &node{}
	n.addStaticChild(c)
	if len(n.staticChildren) != 1 {
		t.Errorf("expected 1 static child, got %d", len(n.staticChildren))
	}
	if n.staticChildren[0] != c {
		t.Errorf("expected %p, got %p", c, n.staticChildren[0])
	}
}
