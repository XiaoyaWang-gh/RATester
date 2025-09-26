package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestInsertFilterChain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	pattern := "/test"
	filterChain := func(next FilterFunc) FilterFunc {
		return func(ctx *context.Context) {
			// pre-processing
			// ...
			next(ctx)
			// post-processing
			// ...
		}
	}

	app.InsertFilterChain(pattern, filterChain)

	// Test if the filter chain is inserted correctly
	_, ok := app.Handlers.routers[pattern]
	if !ok {
		t.Errorf("Expected filter chain for pattern %s to be inserted", pattern)
	}
}
