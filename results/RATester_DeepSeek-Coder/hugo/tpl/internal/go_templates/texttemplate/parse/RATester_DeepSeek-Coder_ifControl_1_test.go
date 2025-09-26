package parse

import (
	"fmt"
	"testing"
)

func TestIfControl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
	}

	node := tree.ifControl()

	if node == nil {
		t.Errorf("Expected node to be not nil")
	}

	if node.Type() != NodeIf {
		t.Errorf("Expected node type to be NodeIf, got %v", node.Type())
	}
}
