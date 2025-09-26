package collections

import (
	"fmt"
	"testing"
)

func TestPush_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stack := &Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Len() != 3 {
		t.Errorf("Expected stack length to be 3, but got %d", stack.Len())
	}

	item, ok := stack.Pop()
	if !ok || item != 3 {
		t.Errorf("Expected to pop 3, but got %d", item)
	}

	if stack.Len() != 2 {
		t.Errorf("Expected stack length to be 2, but got %d", stack.Len())
	}

	item, ok = stack.Peek()
	if !ok || item != 2 {
		t.Errorf("Expected to peek 2, but got %d", item)
	}

	if stack.Len() != 2 {
		t.Errorf("Expected stack length to be 2, but got %d", stack.Len())
	}
}
