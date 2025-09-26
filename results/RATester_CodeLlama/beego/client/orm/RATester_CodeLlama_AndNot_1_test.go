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

	c := Condition{}
	expr := "id = ?"
	args := []interface{}{1}
	c.AndNot(expr, args...)
	if c.params[0].exprs[0] != expr {
		t.Errorf("TestAndNot failed")
	}
	if c.params[0].args[0] != args[0] {
		t.Errorf("TestAndNot failed")
	}
}
