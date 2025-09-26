package math

import (
	"fmt"
	"math"
	"testing"
)

func TestSin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test case 1: Valid input
	result, err := ns.Sin(30)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != math.Sin(30) {
		t.Errorf("Expected %v, but got %v", math.Sin(30), result)
	}

	// Test case 2: Invalid input
	_, err = ns.Sin("abc")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "requires a numeric argument" {
		t.Errorf("Expected error message 'requires a numeric argument', but got '%v'", err.Error())
	}
}
