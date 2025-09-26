package parse

import (
	"fmt"
	"testing"
)

func TestnewList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	node := tree.newList(pos)

	if node.tr != tree {
		t.Errorf("Expected tree to be %v, but got %v", tree, node.tr)
	}

	if node.NodeType != NodeList {
		t.Errorf("Expected NodeType to be %v, but got %v", NodeList, node.NodeType)
	}

	if node.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, node.Pos)
	}
}
