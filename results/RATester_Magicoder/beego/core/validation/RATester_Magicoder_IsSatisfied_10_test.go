package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	maxSize := MaxSize{Max: 5, Key: "test"}

	// Test with a string that is less than or equal to the max size
	result := maxSize.IsSatisfied("hello")
	if !result {
		t.Error("Expected true, got false")
	}

	// Test with a string that is greater than the max size
	result = maxSize.IsSatisfied("hello world")
	if result {
		t.Error("Expected false, got true")
	}

	// Test with a slice that is less than or equal to the max size
	result = maxSize.IsSatisfied([]int{1, 2, 3, 4, 5})
	if !result {
		t.Error("Expected true, got false")
	}

	// Test with a slice that is greater than the max size
	result = maxSize.IsSatisfied([]int{1, 2, 3, 4, 5, 6})
	if result {
		t.Error("Expected false, got true")
	}

	// Test with a non-string and non-slice type
	result = maxSize.IsSatisfied(123)
	if result {
		t.Error("Expected false, got true")
	}
}
