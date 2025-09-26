package collections

import (
	"fmt"
	"testing"
)

func TestLen_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stack := Stack[int]{}

	// Test with an empty stack
	if stack.Len() != 0 {
		t.Errorf("Expected stack length to be 0, but got %d", stack.Len())
	}

	// Test with a stack containing 1 item
	stack.Push(1)
	if stack.Len() != 1 {
		t.Errorf("Expected stack length to be 1, but got %d", stack.Len())
	}

	// Test with a stack containing 5 items
	for i := 2; i <= 5; i++ {
		stack.Push(i)
	}
	if stack.Len() != 5 {
		t.Errorf("Expected stack length to be 5, but got %d", stack.Len())
	}
}
