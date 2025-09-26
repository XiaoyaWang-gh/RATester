package math

import (
	"fmt"
	"math"
	"testing"
)

func TestTan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Tan(45)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result != math.Tan(45) {
		t.Errorf("Expected %v, but got %v", math.Tan(45), result)
	}

	// Test case 2: Invalid input
	_, err = ns.Tan("abc")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
