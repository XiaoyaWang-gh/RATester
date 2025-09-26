package fiber

import (
	"fmt"
	"testing"
)

func TestOnRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	hooks := app.Hooks()

	handler := func(route Route) error {
		// Your test logic here
		return nil
	}

	hooks.OnRoute(handler)

	// Test the OnRoute function
	route := Route{
		Path: "/test",
	}
	hooks.executeOnRouteHooks(route)

	// Add assertions here
}
