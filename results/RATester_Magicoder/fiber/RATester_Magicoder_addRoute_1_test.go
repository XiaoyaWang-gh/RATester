package fiber

import (
	"fmt"
	"testing"
)

func TestaddRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	route := &Route{
		Path: "/test",
	}

	app.addRoute("GET", route)

	if len(app.stack[0]) != 1 {
		t.Errorf("Expected 1 route, got %d", len(app.stack[0]))
	}

	if app.stack[0][0].Path != "/test" {
		t.Errorf("Expected route path to be '/test', got %s", app.stack[0][0].Path)
	}
}
