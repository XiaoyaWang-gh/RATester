package collections

import (
	"fmt"
	"testing"
)

func TestPeek_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item, ok := stack.Peek()
	if !ok {
		t.Fatal("Expected to find an item")
	}
	if item != 3 {
		t.Fatalf("Expected 3, got %d", item)
	}

	stack.Pop()
	item, ok = stack.Peek()
	if !ok {
		t.Fatal("Expected to find an item")
	}
	if item != 2 {
		t.Fatalf("Expected 2, got %d", item)
	}
}
