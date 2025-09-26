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
		t.Errorf("Expected 1 child, got %d", len(n.staticChildren))
	}
}
