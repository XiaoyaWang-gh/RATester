package collections

import (
	"fmt"
	"testing"
)

func TestDrain_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stack := &Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	items := stack.Drain()

	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	if len(stack.items) != 0 {
		t.Errorf("Expected empty stack, got %d items", len(stack.items))
	}
}
