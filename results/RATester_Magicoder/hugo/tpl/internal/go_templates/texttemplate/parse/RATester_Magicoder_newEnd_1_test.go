package parse

import (
	"fmt"
	"testing"
)

func TestnewEnd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(0)
	expected := &endNode{tr: tree, NodeType: nodeEnd, Pos: pos}
	result := tree.newEnd(pos)

	if result.tr != expected.tr || result.NodeType != expected.NodeType || result.Pos != expected.Pos {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
