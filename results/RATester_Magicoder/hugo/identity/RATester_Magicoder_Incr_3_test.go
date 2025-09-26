package identity

import (
	"fmt"
	"testing"
)

func TestIncr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	incr := &IncrementByOne{}

	result := incr.Incr()
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}

	result = incr.Incr()
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}

	result = incr.Incr()
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}
