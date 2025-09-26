package web

import (
	"fmt"
	"testing"
)

func TestExceptMethodAppend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	exceptMethod = []string{}
	ExceptMethodAppend("action1")
	if len(exceptMethod) != 1 {
		t.Errorf("Expected length of exceptMethod to be 1, but got %d", len(exceptMethod))
	}
	if exceptMethod[0] != "action1" {
		t.Errorf("Expected first element of exceptMethod to be 'action1', but got '%s'", exceptMethod[0])
	}
	ExceptMethodAppend("action2")
	if len(exceptMethod) != 2 {
		t.Errorf("Expected length of exceptMethod to be 2, but got %d", len(exceptMethod))
	}
	if exceptMethod[1] != "action2" {
		t.Errorf("Expected second element of exceptMethod to be 'action2', but got '%s'", exceptMethod[1])
	}
}
