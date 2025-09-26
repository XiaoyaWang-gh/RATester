package visitor

import (
	"fmt"
	"testing"
)

func TestCloseListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vm := &Manager{
		listeners: make(map[string]*listenerBundle),
	}
	vm.listeners["test"] = &listenerBundle{}
	vm.CloseListener("test")
	if _, ok := vm.listeners["test"]; ok {
		t.Fatal("CloseListener failed")
	}
}
