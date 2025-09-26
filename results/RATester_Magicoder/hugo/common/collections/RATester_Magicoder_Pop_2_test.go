package collections

import (
	"fmt"
	"testing"
)

func TestPop_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item, ok := stack.Pop()
	if !ok {
		t.Fatal("Expected to get an item")
	}
	if item != 3 {
		t.Fatalf("Expected item to be 3, got %d", item)
	}

	item, ok = stack.Pop()
	if !ok {
		t.Fatal("Expected to get an item")
	}
	if item != 2 {
		t.Fatalf("Expected item to be 2, got %d", item)
	}

	item, ok = stack.Pop()
	if !ok {
		t.Fatal("Expected to get an item")
	}
	if item != 1 {
		t.Fatalf("Expected item to be 1, got %d", item)
	}

	_, ok = stack.Pop()
	if ok {
		t.Fatal("Expected not to get an item")
	}
}
