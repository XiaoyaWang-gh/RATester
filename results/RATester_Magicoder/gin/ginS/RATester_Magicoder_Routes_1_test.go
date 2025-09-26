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
		t.Error("Expected routes to be non-empty")
	}

	for _, route := range routes {
		if route.Method == "" {
			t.Errorf("Expected route method to be set, but got empty string")
		}
		if route.Path == "" {
			t.Errorf("Expected route path to be set, but got empty string")
		}
	}
}
