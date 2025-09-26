package ginS

import (
	"fmt"
	"testing"
)

func TestRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	routes := Routes()

	if len(routes) == 0 {
		t.Errorf("Expected at least one route, got %d", len(routes))
	}

	for _, route := range routes {
		if route.Method == "" {
			t.Errorf("Expected route method to be set, got empty string")
		}

		if route.Path == "" {
			t.Errorf("Expected route path to be set, got empty string")
		}

		if route.Handler == "" {
			t.Errorf("Expected route handler to be set, got empty string")
		}
	}
}
