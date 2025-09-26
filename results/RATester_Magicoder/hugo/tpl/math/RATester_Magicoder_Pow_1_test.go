package math

import (
	"fmt"
	"testing"
)

func TestPow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Pow(2.0, 3.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 8.0 {
		t.Errorf("Expected 8.0, but got %v", result)
	}

	// Test case 2: Invalid input (non-float)
	_, err = ns.Pow("2", 3)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test case 3: Invalid input (non-float)
	_, err = ns.Pow(2, "3")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
