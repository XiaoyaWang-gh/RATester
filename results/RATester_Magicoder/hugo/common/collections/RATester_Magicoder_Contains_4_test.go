package collections

import (
	"fmt"
	"testing"
)

func TestContains_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ss := SortedStringSlice{"apple", "banana", "cherry", "date"}

	// Test case 1: string is in the slice
	if !ss.Contains("banana") {
		t.Error("Expected true, got false")
	}

	// Test case 2: string is not in the slice
	if ss.Contains("grape") {
		t.Error("Expected false, got true")
	}

	// Test case 3: empty string is in the slice
	if !ss.Contains("") {
		t.Error("Expected true, got false")
	}

	// Test case 4: string is in the slice (case insensitive)
	if !ss.Contains("Banana") {
		t.Error("Expected true, got false")
	}

	// Test case 5: string is not in the slice (case insensitive)
	if ss.Contains("Banana") {
		t.Error("Expected false, got true")
	}
}
