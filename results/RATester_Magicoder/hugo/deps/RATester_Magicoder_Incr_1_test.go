package deps

import (
	"fmt"
	"testing"
)

func TestIncr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{
		counter: 0,
	}

	result := b.Incr()

	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	result = b.Incr()

	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}
