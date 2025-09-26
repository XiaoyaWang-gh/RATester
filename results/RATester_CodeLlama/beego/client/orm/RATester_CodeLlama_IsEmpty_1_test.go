package orm

import (
	"fmt"
	"testing"
)

func TestIsEmpty_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Condition{}
	if !c.IsEmpty() {
		t.Errorf("IsEmpty() should return true")
	}
}
