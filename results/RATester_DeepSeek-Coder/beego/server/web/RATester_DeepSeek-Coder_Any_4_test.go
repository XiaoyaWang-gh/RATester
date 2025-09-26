package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestAny_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{},
	}

	testHandler := func(ctx *beecontext.Context) {
		// test logic here
	}

	app.Any("/test", testHandler)

	if app.Handlers.routers["/test"] == nil {
		t.Errorf("Expected /test to be registered")
	}
}
