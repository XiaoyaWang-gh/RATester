package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	// Define the function to be passed to the method under test
	deleteFunc := func(ctx *context.Context) {
		// Implementation of the function
	}

	// Call the method under test
	app.CtrlDelete("/delete", deleteFunc)

	// Assert that the function was added to the ControllerRegister
	_, ok := app.Handlers.routers["/delete"]
	if !ok {
		t.Error("Expected function to be added to ControllerRegister")
	}
}
