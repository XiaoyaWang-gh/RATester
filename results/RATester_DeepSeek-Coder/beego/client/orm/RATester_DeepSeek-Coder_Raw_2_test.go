package orm

import (
	"fmt"
	"testing"
)

func TestRaw_2(t *testing.T) {
	t.Run("TestRaw_EmptySQL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		var c Condition
		c.Raw("", "")
	})

	t.Run("TestRaw_Success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var c Condition
		newC := c.Raw("expr", "sql")
		if newC == &c {
			t.Errorf("Expected new condition to be different from old one")
		}
		if len(c.params) != 1 {
			t.Errorf("Expected 1 param, got %d", len(c.params))
		}
		if c.params[0].exprs[0] != "expr" {
			t.Errorf("Expected exprs[0] to be 'expr', got %s", c.params[0].exprs[0])
		}
		if c.params[0].sql != "sql" {
			t.Errorf("Expected sql to be 'sql', got %s", c.params[0].sql)
		}
		if !c.params[0].isRaw {
			t.Errorf("Expected isRaw to be true")
		}
	})
}
