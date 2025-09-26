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

	s := NewStack[int]()
	if s == nil {
		t.Errorf("NewStack[int]() returned nil")
	}
}
