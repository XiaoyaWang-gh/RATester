package parse

import (
	"fmt"
	"testing"
)

func Testcommand_2(t *testing.T) {
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
					},
				},
			},
		},
	}

	cmd := tree.command()

	if len(cmd.Args) != 1 {
		t.Errorf("Expected 1 argument, got %d", len(cmd.Args))
	}

	variable, ok := cmd.Args[0].(*VariableNode)
	if !ok {
		t.Errorf("Expected VariableNode, got %T", cmd.Args[0])
	}

	if variable.Ident[0] != "var1" {
		t.Errorf("Expected 'var1', got '%s'", variable.Ident[0])
	}
}
