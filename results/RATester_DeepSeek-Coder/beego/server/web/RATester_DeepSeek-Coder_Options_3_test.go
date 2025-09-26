package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestOptions_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	testFunc := func(ctx *beecontext.Context) {
		// test function logic here
	}

	app.Options("/test", testFunc)

	// check if the testFunc is registered as the handler for the "/test" route
	_, ok := app.Handlers.routers["/test"]
	if !ok {
		t.Errorf("Expected testFunc to be registered as the handler for the '/test' route")
	}
}
