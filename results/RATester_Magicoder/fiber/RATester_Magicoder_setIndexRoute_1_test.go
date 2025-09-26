package fiber

import (
	"fmt"
	"testing"
)

func TestsetIndexRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{}
	route := 10
	ctx.setIndexRoute(route)
	if ctx.indexRoute != route {
		t.Errorf("Expected %d, got %d", route, ctx.indexRoute)
	}
}
