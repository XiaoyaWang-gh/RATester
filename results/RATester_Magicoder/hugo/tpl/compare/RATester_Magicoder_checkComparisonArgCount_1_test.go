package compare

import (
	"fmt"
	"testing"
)

func TestcheckComparisonArgCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}

	// Test case with enough arguments
	result := n.checkComparisonArgCount(2, 1, 2)
	if !result {
		t.Errorf("Expected true, got false")
	}

	// Test case with not enough arguments
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	n.checkComparisonArgCount(2)
}
