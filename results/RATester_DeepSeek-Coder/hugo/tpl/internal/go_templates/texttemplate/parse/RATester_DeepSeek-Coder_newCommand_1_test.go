package parse

import (
	"fmt"
	"testing"
)

func TestNewCommand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	commandNode := tree.newCommand(pos)

	if commandNode.tr != tree {
		t.Errorf("Expected tree to be %v, got %v", tree, commandNode.tr)
	}
	if commandNode.NodeType != NodeCommand {
		t.Errorf("Expected NodeType to be %v, got %v", NodeCommand, commandNode.NodeType)
	}
	if commandNode.Pos != pos {
		t.Errorf("Expected Pos to be %v, got %v", pos, commandNode.Pos)
	}
	if len(commandNode.Args) != 0 {
		t.Errorf("Expected Args to be empty, got %v", commandNode.Args)
	}
}
