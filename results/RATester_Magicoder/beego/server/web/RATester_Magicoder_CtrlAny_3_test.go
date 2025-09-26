package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestCtrlAny_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	rootpath := "/any"
	f := func(ctx *context.Context) {
		// test function
	}

	n.CtrlAny(rootpath, f)

	// assert
	if _, ok := n.handlers.routers[rootpath]; !ok {
		t.Errorf("Expected router for path %s, but not found", rootpath)
	}
}
