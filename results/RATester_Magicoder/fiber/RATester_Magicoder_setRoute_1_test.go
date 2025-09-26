package fiber

import (
	"fmt"
	"testing"
)

func TestsetRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{}
	route := &Route{}
	ctx.setRoute(route)

	if ctx.route != route {
		t.Errorf("Expected route to be %v, but got %v", route, ctx.route)
	}
}
