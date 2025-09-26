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

	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	item, ok := s.Pop()
	if !ok {
		t.Fatal("Pop failed")
	}
	if item != 3 {
		t.Errorf("Pop returned %d, want 3", item)
	}
	if s.Len() != 2 {
		t.Errorf("Len() = %d, want 2", s.Len())
	}
	item, ok = s.Pop()
	if !ok {
		t.Fatal("Pop failed")
	}
	if item != 2 {
		t.Errorf("Pop returned %d, want 2", item)
	}
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1", s.Len())
	}
	item, ok = s.Pop()
	if !ok {
		t.Fatal("Pop failed")
	}
	if item != 1 {
		t.Errorf("Pop returned %d, want 1", item)
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}
	item, ok = s.Pop()
	if ok {
		t.Errorf("Pop returned %d, want false", item)
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}
}
