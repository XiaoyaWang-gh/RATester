package doctree

import (
	"fmt"
	"testing"
)

func TestNewSimpleTree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := NewSimpleTree[int]()
	if tree == nil {
		t.Error("NewSimpleTree() should not return nil")
	}
}
