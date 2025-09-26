package context

import (
	"fmt"
	"testing"
)

func TestNewInput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := NewInput()

	if input == nil {
		t.Error("Expected a non-nil BeegoInput, but got nil")
	}

	if len(input.pnames) != 0 {
		t.Errorf("Expected pnames to be empty, but got %v", input.pnames)
	}

	if len(input.pvalues) != 0 {
		t.Errorf("Expected pvalues to be empty, but got %v", input.pvalues)
	}

	if input.data == nil {
		t.Error("Expected data to be a non-nil map, but got nil")
	}
}
