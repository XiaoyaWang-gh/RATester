package orm

import (
	"fmt"
	"testing"
)

func TestSetCond_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{}
	cond := &Condition{}
	qs.SetCond(cond)
	if qs.cond != cond {
		t.Errorf("Expected condition to be set, but it was not.")
	}
}
