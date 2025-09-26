package math

import (
	"fmt"
	"testing"
)

func TestMin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test with multiple inputs
	min, err := ns.Min(1.0, 2.0, 3.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if min != 1.0 {
		t.Errorf("Expected min to be 1.0, but got %f", min)
	}

	// Test with a single input
	min, err = ns.Min(1.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if min != 1.0 {
		t.Errorf("Expected min to be 1.0, but got %f", min)
	}

	// Test with no inputs
	_, err = ns.Min()
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test with non-numeric inputs
	_, err = ns.Min("a", 2.0)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
