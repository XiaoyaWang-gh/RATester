package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()

	handler := func(c Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}

	expectedURL := "/test"
	e.URL(handler, expectedURL)

	routes := e.Routes()
	if len(routes) != 1 {
		t.Errorf("Expected 1 route, got %d", len(routes))
	}

	route := routes[0]
	if route.Path != expectedURL {
		t.Errorf("Expected route path %s, got %s", expectedURL, route.Path)
	}
}
