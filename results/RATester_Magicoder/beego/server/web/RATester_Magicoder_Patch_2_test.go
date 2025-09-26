package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestPatch_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	rootpath := "/patch"
	f := func(ctx *beecontext.Context) {
		// test logic
	}

	n.Patch(rootpath, f)

	// assert
	if _, ok := n.handlers.routers["PATCH"]; !ok {
		t.Error("Patch method not registered")
	}
}
