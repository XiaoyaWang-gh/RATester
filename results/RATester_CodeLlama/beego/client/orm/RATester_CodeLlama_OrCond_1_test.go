package orm

import (
	"fmt"
	"testing"
)

func TestOrCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Condition{}
	cond := &Condition{}
	c.OrCond(cond)
	if len(c.params) != 1 {
		t.Errorf("OrCond failed")
	}
}
