package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	minSize := MinSize{Min: 5}

	// Test case 1: Validate a string with length greater than or equal to Min
	str := "Hello"
	if !minSize.IsSatisfied(str) {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Validate a string with length less than Min
	str = "Hi"
	if minSize.IsSatisfied(str) {
		t.Errorf("Expected false, but got true")
	}

	// Test case 3: Validate a slice with length greater than or equal to Min
	slice := []int{1, 2, 3, 4, 5}
	if !minSize.IsSatisfied(slice) {
		t.Errorf("Expected true, but got false")
	}

	// Test case 4: Validate a slice with length less than Min
	slice = []int{1, 2, 3, 4}
	if minSize.IsSatisfied(slice) {
		t.Errorf("Expected false, but got true")
	}

	// Test case 5: Validate a non-string and non-slice type
	num := 10
	if minSize.IsSatisfied(num) {
		t.Errorf("Expected false, but got true")
	}
}
