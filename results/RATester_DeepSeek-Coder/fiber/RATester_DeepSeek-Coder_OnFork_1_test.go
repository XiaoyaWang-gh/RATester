package fiber

import (
	"fmt"
	"testing"
)

func TestOnFork_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	hooks := app.Hooks()

	handlerCalled := false
	handler := func(pid int) error {
		handlerCalled = true
		return nil
	}

	hooks.OnFork(handler)
	hooks.executeOnForkHooks(1234)

	if !handlerCalled {
		t.Errorf("Expected handler to be called")
	}
}
