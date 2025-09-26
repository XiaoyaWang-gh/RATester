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

	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Len() != 3 {
		t.Errorf("expected 3, got %d", s.Len())
	}
}
