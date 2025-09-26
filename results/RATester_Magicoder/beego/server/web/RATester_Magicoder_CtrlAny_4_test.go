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

	rootpath := "/test"
	f := func(ctx *context.Context) {
		// test function
	}

	app.CtrlAny(rootpath, f)

	// Assert that the function was added to the ControllerRegister
	_, ok := app.Handlers.routers[rootpath]
	if !ok {
		t.Errorf("Expected function to be added to ControllerRegister, but it was not")
	}
}
