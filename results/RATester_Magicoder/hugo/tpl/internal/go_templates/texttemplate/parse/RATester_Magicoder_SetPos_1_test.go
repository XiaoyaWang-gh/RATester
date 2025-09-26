package parse

import (
	"fmt"
	"testing"
)

func TestSetPos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	node := &IdentifierNode{
		Pos:   10,
		Ident: "test",
	}

	newPos := Pos(20)
	node.SetPos(newPos)

	if node.Pos != newPos {
		t.Errorf("Expected Pos to be %d, but got %d", newPos, node.Pos)
	}
}
