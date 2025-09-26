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

	cond := &Condition{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cond.AndNot("", nil)
}
