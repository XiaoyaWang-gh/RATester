package parse

import (
	"fmt"
	"testing"
)

func TestnewChain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(0)
	node := &ChainNode{}
	chainNode := tree.newChain(pos, node)

	if chainNode.tr != tree {
		t.Errorf("Expected tree to be %v, but got %v", tree, chainNode.tr)
	}

	if chainNode.NodeType != NodeChain {
		t.Errorf("Expected NodeType to be %v, but got %v", NodeChain, chainNode.NodeType)
	}

	if chainNode.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, chainNode.Pos)
	}

	if chainNode.Node != node {
		t.Errorf("Expected Node to be %v, but got %v", node, chainNode.Node)
	}
}
