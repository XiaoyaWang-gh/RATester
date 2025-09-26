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

	expected := 1
	result := b.Incr()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	expected = 2
	result = b.Incr()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
