package math

import (
	"fmt"
	"math"
	"testing"
)

func TestAtan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Atan(1.0)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != math.Atan(1.0) {
		t.Errorf("Expected %v, but got %v", math.Atan(1.0), result)
	}

	// Test case 2: Invalid input
	_, err = ns.Atan("invalid")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
