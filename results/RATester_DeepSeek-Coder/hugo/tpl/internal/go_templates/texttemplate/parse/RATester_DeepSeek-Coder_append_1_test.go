package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend_1(t *testing.T) {
	testCases := []struct {
		name     string
		original *CommandNode
		arg      Node
		expected *CommandNode
	}{
		{
			name: "append to empty CommandNode",
			original: &CommandNode{
				Args: []Node{},
			},
			arg: &IdentifierNode{
				Ident: "test",
			},
			expected: &CommandNode{
				Args: []Node{
					&IdentifierNode{
						Ident: "test",
					},
				},
			},
		},
		{
			name: "append to non-empty CommandNode",
			original: &CommandNode{
				Args: []Node{
					&IdentifierNode{
						Ident: "existing",
					},
				},
			},
			arg: &IdentifierNode{
				Ident: "new",
			},
			expected: &CommandNode{
				Args: []Node{
					&IdentifierNode{
						Ident: "existing",
					},
					&IdentifierNode{
						Ident: "new",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.original.append(tc.arg)
			if !reflect.DeepEqual(tc.original, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.original)
			}
		})
	}
}
