package parse

import (
	"fmt"
	"testing"
)

func TestendControl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
		Root: &ListNode{
			Nodes: []Node{},
		},
	}
	expected := &endNode{
		Pos: 0,
		tr:  tree,
	}
	result := tree.endControl()
	if result.Type() != expected.Type() {
		t.Errorf("Expected type %v, but got %v", expected.Type(), result.Type())
	}
	if result.Position() != expected.Position() {
		t.Errorf("Expected position %v, but got %v", expected.Position(), result.Position())
	}
}
