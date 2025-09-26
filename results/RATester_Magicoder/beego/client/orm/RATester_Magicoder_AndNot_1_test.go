package orm

import (
	"fmt"
	"testing"
)

func TestAndNot_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cond := Condition{}

	// Test with empty expr
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cond.AndNot("", "arg1")

	// Test with empty args
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cond.AndNot("expr", "")

	// Test with valid input
	newCond := cond.AndNot("expr", "arg1")
	if newCond == nil {
		t.Errorf("Expected a new condition, but got nil")
	}
}
