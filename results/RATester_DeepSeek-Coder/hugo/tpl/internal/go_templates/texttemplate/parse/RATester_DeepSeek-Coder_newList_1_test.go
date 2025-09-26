package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	expected := &ListNode{
		NodeType: NodeList,
		Pos:      pos,
		tr:       tree,
		Nodes:    make([]Node, 0),
	}
	result := tree.newList(pos)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
