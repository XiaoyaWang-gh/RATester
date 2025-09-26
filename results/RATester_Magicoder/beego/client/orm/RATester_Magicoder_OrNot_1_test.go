package orm

import (
	"fmt"
	"testing"
)

func TestOrNot_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cond := Condition{}

	// Test case 1: Normal case
	expr := "name = ?"
	args := []interface{}{"John"}
	cond.OrNot(expr, args...)
	if len(cond.params) != 1 {
		t.Errorf("Expected 1 param, got %d", len(cond.params))
	}

	// Test case 2: Empty expr
	expr = ""
	args = []interface{}{"John"}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	cond.OrNot(expr, args...)

	// Test case 3: Empty args
	expr = "name = ?"
	args = []interface{}{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	cond.OrNot(expr, args...)
}
