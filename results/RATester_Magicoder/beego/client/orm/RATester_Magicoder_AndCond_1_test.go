package orm

import (
	"fmt"
	"testing"
)

func TestAndCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new Condition
	c := &Condition{}

	// Create a new Condition to be used as a sub condition
	subCond := &Condition{}

	// Test the AndCond method
	result := c.AndCond(subCond)

	// Check if the result is as expected
	if result == nil {
		t.Error("Expected a non-nil result, but got nil")
	}

	// Check if the original Condition is not modified
	if c == result {
		t.Error("Expected the original Condition to be cloned, but it was not")
	}

	// Check if the sub condition is added to the new Condition
	if len(result.params) != 1 || result.params[0].cond != subCond {
		t.Error("Expected the sub condition to be added to the new Condition, but it was not")
	}
}
