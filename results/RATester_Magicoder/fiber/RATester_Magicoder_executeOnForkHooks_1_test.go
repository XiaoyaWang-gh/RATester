package fiber

import (
	"fmt"
	"testing"
)

func TestexecuteOnForkHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hooks := &Hooks{}

	pid := 12345

	handlerCalled := false
	handler := func(pid int) error {
		handlerCalled = true
		if pid != pid {
			t.Errorf("Expected pid %d, but got %d", pid, pid)
		}
		return nil
	}

	hooks.OnFork(handler)
	hooks.executeOnForkHooks(pid)

	if !handlerCalled {
		t.Error("Handler was not called")
	}
}
