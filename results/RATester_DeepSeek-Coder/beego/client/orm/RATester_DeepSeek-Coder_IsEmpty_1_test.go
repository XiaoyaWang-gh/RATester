package orm

import (
	"fmt"
	"testing"
)

func TestIsEmpty_1(t *testing.T) {
	t.Run("TestIsEmpty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cond := &Condition{params: []condValue{}}
		if !cond.IsEmpty() {
			t.Errorf("Expected IsEmpty to return true, got false")
		}

		cond = &Condition{params: []condValue{{exprs: []string{"test"}}}}
		if cond.IsEmpty() {
			t.Errorf("Expected IsEmpty to return false, got true")
		}
	})
}
