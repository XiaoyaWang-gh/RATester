package fiber

import (
	"fmt"
	"testing"
)

func TestOnShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	hooks := app.Hooks()

	handlerCalled := false
	handler := func() error {
		handlerCalled = true
		return nil
	}

	hooks.OnShutdown(handler)

	hooks.executeOnShutdownHooks()

	if !handlerCalled {
		t.Errorf("Expected handler to be called")
	}
}
