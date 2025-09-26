package fiber

import (
	"fmt"
	"testing"
)

func TestOnName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	hooks := app.Hooks()

	handler := func(route Route) error {
		// Test logic here
		return nil
	}

	hooks.OnName(handler)

	if len(hooks.onName) != 1 {
		t.Errorf("Expected 1 handler, got %d", len(hooks.onName))
	}

	hooks.executeOnNameHooks(Route{})

	// Add more test cases as needed
}
