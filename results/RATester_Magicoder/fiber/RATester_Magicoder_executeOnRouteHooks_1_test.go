package fiber

import (
	"fmt"
	"testing"
)

func TestexecuteOnRouteHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hooks := &Hooks{
		app: &App{
			mountFields: &mountFields{
				mountPath: "/test",
			},
		},
	}

	route := &Route{
		path: "/test",
	}

	err := hooks.executeOnRouteHooks(*route)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if route.path != "/test/test" {
		t.Errorf("Expected /test/test, got %v", route.path)
	}
}
