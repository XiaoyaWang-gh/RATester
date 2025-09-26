package orm

import (
	"fmt"
	"testing"
)

func TestOrNotCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cond1 := &Condition{}
	cond2 := &Condition{}

	// Test case 1: Adding a condition to the existing condition
	cond1.OrNotCond(cond2)
	if len(cond1.params) != 1 {
		t.Errorf("Expected 1 condition, got %d", len(cond1.params))
	}
	if cond1.params[0].cond != cond2 {
		t.Errorf("Expected condition %v, got %v", cond2, cond1.params[0].cond)
	}
	if cond1.params[0].isNot != true {
		t.Errorf("Expected isNot to be true, got false")
	}
	if cond1.params[0].isOr != true {
		t.Errorf("Expected isOr to be true, got false")
	}

	// Test case 2: Adding nil condition to the existing condition
	cond1.OrNotCond(nil)
	if len(cond1.params) != 1 {
		t.Errorf("Expected 1 condition, got %d", len(cond1.params))
	}

	// Test case 3: Adding the same condition to the existing condition
	cond1.OrNotCond(cond1)
	if len(cond1.params) != 1 {
		t.Errorf("Expected 1 condition, got %d", len(cond1.params))
	}
}
