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

	rootpath := "/put"
	f := func(ctx *context.Context) {
		// test function
	}

	n.CtrlPut(rootpath, f)

	// assert
	if _, ok := n.handlers.routers[rootpath]; !ok {
		t.Errorf("Expected router for %s, but not found", rootpath)
	}
}
