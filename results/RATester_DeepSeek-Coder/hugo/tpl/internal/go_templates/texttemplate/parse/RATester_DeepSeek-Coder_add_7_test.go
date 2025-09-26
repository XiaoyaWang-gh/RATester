package parse

import (
	"fmt"
	"testing"
)

func TestAdd_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{
		Name:    "test",
		treeSet: make(map[string]*Tree),
	}
	tree.add()
	if _, ok := tree.treeSet["test"]; !ok {
		t.Errorf("Expected tree to be added to treeSet")
	}
}
