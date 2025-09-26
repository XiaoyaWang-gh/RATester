package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlPut_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	testFunc := func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	}

	n.CtrlPut("/put", testFunc)

	if n.handlers.routers["/test/put"] == nil {
		t.Errorf("Expected /test/put to be registered")
	}
}
