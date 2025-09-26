package parse

import (
	"fmt"
	"testing"
)

func TestnewDot_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(0)
	dotNode := tree.newDot(pos)

	if dotNode.tr != tree {
		t.Errorf("Expected tree to be %v, but got %v", tree, dotNode.tr)
	}

	if dotNode.NodeType != NodeDot {
		t.Errorf("Expected NodeType to be %v, but got %v", NodeDot, dotNode.NodeType)
	}

	if dotNode.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, dotNode.Pos)
	}
}
