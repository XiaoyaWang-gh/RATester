package fiber

import (
	"fmt"
	"testing"
)

func TestCopyRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := &App{}
	route := &Route{
		use: true,
	}

	copiedRoute := app.copyRoute(route)

	if copiedRoute.use != route.use {
		t.Errorf("Expected copiedRoute.use to be %v, got %v", route.use, copiedRoute.use)
	}
}
