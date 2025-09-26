package orm

import (
	"fmt"
	"testing"
)

func TestAndNotCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cond1 := &Condition{}
	cond2 := &Condition{}

	result := cond1.AndNotCond(cond2)

	if result == cond1 {
		t.Errorf("Expected a new condition, got the same condition")
	}

	if len(result.params) != 1 {
		t.Errorf("Expected 1 condition, got %d", len(result.params))
	}

	param := result.params[0]
	if param.cond != cond2 {
		t.Errorf("Expected condition %v, got %v", cond2, param.cond)
	}

	if !param.isCond {
		t.Errorf("Expected isCond to be true")
	}

	if !param.isNot {
		t.Errorf("Expected isNot to be true")
	}
}
