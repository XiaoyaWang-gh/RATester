package collections

import (
	"fmt"
	"testing"
)

func TestPeek_3(t *testing.T) {
	t.Parallel()

	t.Run("empty stack", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		s := &Stack[int]{}
		v, ok := s.Peek()
		if ok {
			t.Errorf("expected ok to be false, got true")
		}
		if v != 0 {
			t.Errorf("expected v to be 0, got %v", v)
		}
	})

	t.Run("non-empty stack", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		s := &Stack[int]{items: []int{1, 2, 3}}
		v, ok := s.Peek()
		if !ok {
			t.Errorf("expected ok to be true, got false")
		}
		if v != 3 {
			t.Errorf("expected v to be 3, got %v", v)
		}
	})
}
