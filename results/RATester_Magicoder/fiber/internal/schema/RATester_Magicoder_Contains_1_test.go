package schema

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	options := tagOptions{"option1", "option2", "option3"}

	// Test case 1: option exists
	if !options.Contains("option1") {
		t.Error("Expected to find option1")
	}

	// Test case 2: option does not exist
	if options.Contains("option4") {
		t.Error("Expected not to find option4")
	}

	// Test case 3: empty option
	if options.Contains("") {
		t.Error("Expected not to find empty option")
	}
}
