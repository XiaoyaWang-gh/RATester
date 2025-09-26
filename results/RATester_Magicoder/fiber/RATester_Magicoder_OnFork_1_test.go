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

	handler := func(pid int) error {
		// Your test logic here
		return nil
	}

	hooks.OnFork(handler)

	// Test the behavior of the OnFork method
	// For example, you can call the executeOnForkHooks method and check if the handler was called correctly
	hooks.executeOnForkHooks(123)

	// Add more test cases as needed
}
