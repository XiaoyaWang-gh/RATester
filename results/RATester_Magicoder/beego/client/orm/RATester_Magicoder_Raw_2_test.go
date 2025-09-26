package orm

import (
	"fmt"
	"testing"
)

func TestRaw_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Condition{}

	sql := "SELECT * FROM table WHERE id = ?"
	expr := "id = ?"

	result := c.Raw(expr, sql)

	if len(result.params) != 1 {
		t.Errorf("Expected 1 param, got %d", len(result.params))
	}

	if result.params[0].sql != sql {
		t.Errorf("Expected sql to be %s, got %s", sql, result.params[0].sql)
	}

	if result.params[0].isRaw != true {
		t.Errorf("Expected isRaw to be true, got false")
	}

	if result.params[0].exprs[0] != expr {
		t.Errorf("Expected exprs[0] to be %s, got %s", expr, result.params[0].exprs[0])
	}
}
