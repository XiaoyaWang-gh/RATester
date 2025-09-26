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

	condition := &Condition{params: []condValue{}}
	if !condition.IsEmpty() {
		t.Error("Expected IsEmpty to return true, but got false")
	}

	condition = &Condition{params: []condValue{{exprs: []string{"test"}}}}
	if condition.IsEmpty() {
		t.Error("Expected IsEmpty to return false, but got true")
	}
}
