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

	// Define a test handler
	testHandler := func() error {
		// Test logic here
		return nil
	}

	// Add the test handler to the hooks
	hooks.OnShutdown(testHandler)

	// Execute the hooks
	hooks.executeOnShutdownHooks()

	// Add assertions here to verify the testHandler was called
}
