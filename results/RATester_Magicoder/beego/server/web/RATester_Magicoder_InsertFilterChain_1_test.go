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
			// do something before
			next(ctx)
			// do something after
		}
	}

	app.InsertFilterChain(pattern, filterChain)

	// assert something
}
