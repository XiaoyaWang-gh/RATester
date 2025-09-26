package math

import (
	"fmt"
	"math"
	"testing"
)

func TestAtan2_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Atan2(1, 1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != math.Pi/4 {
		t.Errorf("Expected %v, but got %v", math.Pi/4, result)
	}

	// Test case 2: Invalid input (non-numeric)
	_, err = ns.Atan2("a", 1)
	if err == nil {
		t.Error("Expected an error, but got none")
	}

	// Test case 3: Invalid input (non-numeric)
	_, err = ns.Atan2(1, "b")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
