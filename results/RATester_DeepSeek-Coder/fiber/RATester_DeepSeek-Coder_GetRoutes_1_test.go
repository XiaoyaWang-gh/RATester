package fiber

import (
	"fmt"
	"testing"
)

func TestGetRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := new(App)
	route1 := &Route{Method: "GET", Path: "/test1"}
	route2 := &Route{Method: "POST", Path: "/test2"}
	app.stack = append(app.stack, []*Route{route1, route2})

	routes := app.GetRoutes()
	if len(routes) != 2 {
		t.Errorf("Expected 2 routes, got %d", len(routes))
	}

	routes = app.GetRoutes(true)
	if len(routes) != 1 {
		t.Errorf("Expected 1 route, got %d", len(routes))
	}
}
