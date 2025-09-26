package orm

import (
	"fmt"
	"testing"
)

func TestOrCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new condition
	cond1 := &Condition{}
	cond2 := &Condition{}

	// Test the OrCond method
	result := cond1.OrCond(cond2)

	// Assert that the result is not nil
	if result == nil {
		t.Error("Expected a non-nil result, but got nil")
	}

	// Assert that the result is not the same as cond1 or cond2
	if result == cond1 || result == cond2 {
		t.Error("Expected a new condition, but got the same condition")
	}

	// Assert that the result has the correct parameters
	if len(result.params) != 1 {
		t.Errorf("Expected 1 parameter, but got %d", len(result.params))
	}

	param := result.params[0]
	if param.cond != cond2 {
		t.Error("Expected the second condition, but got a different condition")
	}
	if !param.isCond {
		t.Error("Expected isCond to be true")
	}
	if !param.isOr {
		t.Error("Expected isOr to be true")
	}
}
