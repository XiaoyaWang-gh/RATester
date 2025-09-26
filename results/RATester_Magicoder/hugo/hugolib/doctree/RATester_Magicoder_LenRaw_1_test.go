package doctree

import (
	"fmt"
	"testing"
)

func TestLenRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &TreeShiftTree[int]{}

	// Add some data to the tree
	tree.Insert("key1", 1)
	tree.Insert("key2", 2)
	tree.Insert("key3", 3)

	// Test the LenRaw method
	expected := 6
	actual := tree.LenRaw()
	if actual != expected {
		t.Errorf("Expected LenRaw to return %d, but got %d", expected, actual)
	}
}
