package math

import (
	"fmt"
	"testing"
)

func TestSqrt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	// Test with valid input
	result, err := ns.Sqrt(4)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result != 2 {
		t.Errorf("Expected 2, but got %v", result)
	}

	// Test with invalid input
	_, err = ns.Sqrt("a")
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	if err.Error() != "Sqrt operator can't be used with non integer or float value" {
		t.Errorf("Expected 'Sqrt operator can't be used with non integer or float value', but got %v", err.Error())
	}
}
