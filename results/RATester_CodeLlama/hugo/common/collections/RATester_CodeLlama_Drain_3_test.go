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

	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	items := s.Drain()
	if len(items) != 3 {
		t.Errorf("expected 3 items, got %d", len(items))
	}
	if items[0] != 1 || items[1] != 2 || items[2] != 3 {
		t.Errorf("expected [1, 2, 3], got %v", items)
	}
	if s.Len() != 0 {
		t.Errorf("expected empty stack, got %d items", s.Len())
	}
}
