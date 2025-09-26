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

	var c Condition
	expr := "id = ?"
	args := []interface{}{1}
	c.OrNot(expr, args...)
	if c.IsEmpty() {
		t.Errorf("OrNot failed")
	}
}
