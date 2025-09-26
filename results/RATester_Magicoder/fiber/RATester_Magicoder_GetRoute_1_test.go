package fiber

import (
	"fmt"
	"testing"
)

func TestGetRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		stack: [][]*Route{
			{
				{
					Name: "testRoute",
				},
			},
		},
	}

	route := app.GetRoute("testRoute")

	if route.Name != "testRoute" {
		t.Errorf("Expected route name to be 'testRoute', but got '%s'", route.Name)
	}
}
