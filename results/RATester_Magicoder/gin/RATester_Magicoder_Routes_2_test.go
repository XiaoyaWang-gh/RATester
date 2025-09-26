package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRoutes_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		// Initialize your Engine instance here
	}

	expectedRoutes := RoutesInfo{
		// Define the expected routes here
	}

	routes := engine.Routes()

	if !reflect.DeepEqual(routes, expectedRoutes) {
		t.Errorf("Expected routes: %v, but got: %v", expectedRoutes, routes)
	}
}
