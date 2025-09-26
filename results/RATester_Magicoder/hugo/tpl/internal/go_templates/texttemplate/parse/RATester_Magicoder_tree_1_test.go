package parse

import (
	"fmt"
	"testing"
)

func Testtree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	node := &TemplateNode{
		tr: &Tree{
			Name: "test",
		},
	}

	tree := node.tree()

	if tree.Name != "test" {
		t.Errorf("Expected tree name to be 'test', but got '%s'", tree.Name)
	}
}
