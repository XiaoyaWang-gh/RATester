package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlPut_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	testFunc := func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	}

	app.CtrlPut("/test", testFunc)

	if _, ok := app.Handlers.routers["/test"]; !ok {
		t.Errorf("Expected /test to be registered in Handlers.routers")
	}
}
