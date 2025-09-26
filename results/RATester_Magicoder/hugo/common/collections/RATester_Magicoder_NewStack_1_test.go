package collections

import (
	"fmt"
	"testing"
)

func TestNewStack_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stack := NewStack[int]()

	if stack == nil {
		t.Error("NewStack() should not return nil")
	}

	if stack.Len() != 0 {
		t.Error("NewStack() should return an empty stack")
	}
}
