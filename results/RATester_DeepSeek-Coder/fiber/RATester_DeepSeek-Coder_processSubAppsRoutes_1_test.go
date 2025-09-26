package fiber

import (
	"fmt"
	"testing"
)

func TestProcessSubAppsRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		mountFields: &mountFields{
			appList: make(map[string]*App),
		},
	}

	// Add a sub-app to the app list
	subApp := &App{
		mountFields: &mountFields{
			appList: make(map[string]*App),
		},
	}
	app.mountFields.appList["subApp"] = subApp

	// Add a route to the sub-app
	subApp.addRoute("GET", &Route{
		Method: "GET",
		Path:   "/test",
	})

	// Call the function to process the sub-apps routes
	app.processSubAppsRoutes()

	// Check if the route was added to the main app's stack
	if len(app.stack) == 0 {
		t.Errorf("Expected routes to be processed, but the stack is empty")
	}

	// Check if the route was added to the correct position in the stack
	if app.stack[0][0].Path != "/test" {
		t.Errorf("Expected route to be at position 0, but it was at position %d", app.stack[0][0].pos)
	}
}
