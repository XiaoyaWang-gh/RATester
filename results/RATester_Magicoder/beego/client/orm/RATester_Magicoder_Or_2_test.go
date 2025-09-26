package orm

import (
	"fmt"
	"testing"
)

func TestOr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cond := Condition{}

	// Test case 1: Valid input
	expr := "column1 = ? AND column2 = ?"
	args := []interface{}{1, 2}
	newCond := cond.Or(expr, args...)
	if newCond == nil {
		t.Error("Expected non-nil condition, but got nil")
	}

	// Test case 2: Empty expr
	expr = ""
	args = []interface{}{1, 2}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but got none")
		}
	}()
	cond.Or(expr, args...)

	// Test case 3: Empty args
	expr = "column1 = ?"
	args = []interface{}{}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but got none")
		}
	}()
	cond.Or(expr, args...)
}
