package doctree

import (
	"fmt"
	"testing"
)

func TestWalkPrefixRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &TreeShiftTree[int]{}

	// Define a test function that returns true to continue walking
	testFunc := func(s string, v int) (bool, error) {
		// Perform some operation on the key-value pair
		// For example, print the key and value
		fmt.Printf("Key: %s, Value: %d\n", s, v)
		return true, nil
	}

	// Call the method under test
	err := tree.WalkPrefixRaw(LockType(0), "prefix", testFunc)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
