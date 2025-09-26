package parse

import (
	"fmt"
	"testing"
)

func TestnewBreak_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	line := 1
	node := tree.newBreak(pos, line)

	if node.tr != tree {
		t.Errorf("Expected tree to be %v, but got %v", tree, node.tr)
	}

	if node.NodeType != NodeBreak {
		t.Errorf("Expected NodeType to be %v, but got %v", NodeBreak, node.NodeType)
	}

	if node.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, node.Pos)
	}

	if node.Line != line {
		t.Errorf("Expected Line to be %v, but got %v", line, node.Line)
	}
}
