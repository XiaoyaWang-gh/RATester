package collections

import (
	"fmt"
	"testing"
)

func TestDrainMatching_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stack := &Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	result := stack.DrainMatching(func(i int) bool {
		return i%2 == 0
	})

	if len(result) != 2 {
		t.Errorf("Expected 2 items, got %d", len(result))
	}

	if result[0] != 4 || result[1] != 2 {
		t.Errorf("Unexpected items in result: %v", result)
	}

	if stack.Len() != 3 {
		t.Errorf("Expected stack to have 3 items, got %d", stack.Len())
	}
}
