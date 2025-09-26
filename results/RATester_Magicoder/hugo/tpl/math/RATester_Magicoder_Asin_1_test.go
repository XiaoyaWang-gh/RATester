package math

import (
	"fmt"
	"math"
	"testing"
)

func TestAsin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Asin(1.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != math.Asin(1.0) {
		t.Errorf("Expected %f, but got %f", math.Asin(1.0), result)
	}

	// Test case 2: Invalid input (non-numeric)
	_, err = ns.Asin("invalid")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	// Test case 3: Invalid input (nil)
	_, err = ns.Asin(nil)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
