package parse

import (
	"fmt"
	"testing"
)

func TestCommand_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
		Root: &ListNode{
			Nodes: []Node{
				&CommandNode{
					Args: []Node{
						&VariableNode{
							Ident: []string{"var1"},
						},
						&VariableNode{
							Ident: []string{"var2"},
						},
					},
				},
			},
		},
	}

	cmd := tree.command()
	if len(cmd.Args) != 2 {
		t.Errorf("Expected 2 arguments, got %d", len(cmd.Args))
	}

	_, ok := cmd.Args[0].(*VariableNode)
	if !ok {
		t.Errorf("Expected first argument to be a VariableNode, got %T", cmd.Args[0])
	}

	_, ok = cmd.Args[1].(*VariableNode)
	if !ok {
		t.Errorf("Expected second argument to be a VariableNode, got %T", cmd.Args[1])
	}
}
