package math

import (
	"fmt"
	"math"
	"testing"
)

func TestCos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Cos(3.14)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != math.Cos(3.14) {
		t.Errorf("Expected %v, but got %v", math.Cos(3.14), result)
	}

	// Test case 2: Invalid input
	_, err = ns.Cos("invalid")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// Test case 3: Zero input
	result, err = ns.Cos(0)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != math.Cos(0) {
		t.Errorf("Expected %v, but got %v", math.Cos(0), result)
	}
}
