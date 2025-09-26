package parse

import (
	"fmt"
	"testing"
)

func TestnewAction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	line := 1
	pipe := &PipeNode{}
	action := tree.newAction(pos, line, pipe)

	if action.tr != tree {
		t.Errorf("Expected tree to be %v, but got %v", tree, action.tr)
	}
	if action.NodeType != NodeAction {
		t.Errorf("Expected NodeType to be %v, but got %v", NodeAction, action.NodeType)
	}
	if action.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, action.Pos)
	}
	if action.Line != line {
		t.Errorf("Expected Line to be %v, but got %v", line, action.Line)
	}
	if action.Pipe != pipe {
		t.Errorf("Expected Pipe to be %v, but got %v", pipe, action.Pipe)
	}
}
