package visitor

import (
	"fmt"
	"sync"
	"testing"
)

func TestCloseListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	vm := &Manager{
		listeners: make(map[string]*listenerBundle),
		mu:        sync.RWMutex{},
	}

	vm.listeners["test"] = &listenerBundle{}

	vm.CloseListener("test")

	if _, ok := vm.listeners["test"]; ok {
		t.Errorf("Expected listener to be deleted")
	}
}
