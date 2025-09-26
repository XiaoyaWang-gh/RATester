package parse

import (
	"fmt"
	"testing"
)

func Testadd_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name: "test",
		treeSet: map[string]*Tree{
			"test": &Tree{
				Root: &ListNode{
					Nodes: []Node{},
				},
			},
		},
	}

	tree.add()

	if tree.treeSet["test"].Root != nil {
		t.Errorf("Expected nil root, got %v", tree.treeSet["test"].Root)
	}
}
