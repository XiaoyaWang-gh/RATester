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
		t.Errorf("NewInput() = %v, want %v", input, "not nil")
	}
}
