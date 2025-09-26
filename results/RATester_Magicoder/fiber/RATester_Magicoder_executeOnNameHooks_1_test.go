package fiber

import (
	"fmt"
	"testing"
)

func TestexecuteOnNameHooks_1(t *testing.T) {
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
		onName: []func(Route) error{
			func(route Route) error {
				route.path = "/test" + route.path
				return nil
			},
		},
	}

	route := Route{
		path: "/test",
	}

	err := hooks.executeOnNameHooks(route)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if route.path != "/test/test" {
		t.Errorf("Expected path to be /test/test, but got %v", route.path)
	}
}
