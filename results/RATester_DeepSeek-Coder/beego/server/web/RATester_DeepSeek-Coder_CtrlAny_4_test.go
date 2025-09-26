package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlAny_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	testFunc := func(ctx *context.Context) {
		// test function implementation
	}

	app.CtrlAny("/test", testFunc)

	// check if the testFunc is added to the router
	_, ok := app.Handlers.routers["/test"]
	if !ok {
		t.Errorf("Expected testFunc to be added to the router, but it was not")
	}
}
