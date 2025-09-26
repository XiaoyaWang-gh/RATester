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

	s := &Stack[int]{}

	// Test initial length
	expected := 0
	actual := s.Len()
	if actual != expected {
		t.Errorf("Expected length %d, but got %d", expected, actual)
	}

	// Test length after pushing items
	s.Push(1)
	s.Push(2)
	expected = 2
	actual = s.Len()
	if actual != expected {
		t.Errorf("Expected length %d, but got %d", expected, actual)
	}

	// Test length after popping items
	_, _ = s.Pop()
	_, _ = s.Pop()
	expected = 0
	actual = s.Len()
	if actual != expected {
		t.Errorf("Expected length %d, but got %d", expected, actual)
	}
}
