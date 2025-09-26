package math

import (
	"fmt"
	"testing"
)

func TestSum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Sum of two positive numbers
	sum, err := ns.Sum(3.5, 2.5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sum != 6.0 {
		t.Errorf("Expected 6.0, but got %v", sum)
	}

	// Test case 2: Sum of a positive and a negative number
	sum, err = ns.Sum(3.5, -2.5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sum != 1.0 {
		t.Errorf("Expected 1.0, but got %v", sum)
	}

	// Test case 3: Sum of two zeroes
	sum, err = ns.Sum(0.0, 0.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sum != 0.0 {
		t.Errorf("Expected 0.0, but got %v", sum)
	}

	// Test case 4: Sum of a positive and a zero
	sum, err = ns.Sum(3.5, 0.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sum != 3.5 {
		t.Errorf("Expected 3.5, but got %v", sum)
	}

	// Test case 5: Sum of a negative and a zero
	sum, err = ns.Sum(-3.5, 0.0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sum != -3.5 {
		t.Errorf("Expected -3.5, but got %v", sum)
	}
}
