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
	c.Raw("a = ?", "1")
	if len(c.params) != 1 {
		t.Errorf("c.params length is not 1")
	}
	if c.params[0].exprs[0] != "a = ?" {
		t.Errorf("c.params[0].exprs[0] is not a = ?")
	}
	if c.params[0].sql != "1" {
		t.Errorf("c.params[0].sql is not 1")
	}
	if c.params[0].isRaw != true {
		t.Errorf("c.params[0].isRaw is not true")
	}
}
