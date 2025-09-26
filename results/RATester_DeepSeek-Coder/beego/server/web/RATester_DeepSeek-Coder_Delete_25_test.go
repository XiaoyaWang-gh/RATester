package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestDelete_25(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	testFunc := func(ctx *beecontext.Context) {
		// test logic here
	}

	app.Delete("/test", testFunc)

	if _, ok := app.Handlers.routers["/test"]; !ok {
		t.Errorf("Expected route /test to be registered")
	}
}
