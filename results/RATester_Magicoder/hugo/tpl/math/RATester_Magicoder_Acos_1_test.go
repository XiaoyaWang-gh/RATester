package math

import (
	"fmt"
	"math"
	"testing"
)

func TestAcos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Acos(0.5)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != math.Acos(0.5) {
		t.Errorf("Expected %v, but got %v", math.Acos(0.5), result)
	}

	// Test case 2: Invalid input (non-numeric)
	_, err = ns.Acos("invalid")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
