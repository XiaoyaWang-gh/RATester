package orm

import (
	"fmt"
	"testing"
)

func TestGetCond_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{
		cond: &Condition{},
	}

	cond := qs.GetCond()

	if cond != qs.cond {
		t.Errorf("Expected %v, got %v", qs.cond, cond)
	}
}
