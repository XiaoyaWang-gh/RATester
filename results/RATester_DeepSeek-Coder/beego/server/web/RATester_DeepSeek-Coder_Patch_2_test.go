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

	called := false
	f := func(ctx *beecontext.Context) {
		called = true
	}

	n.Patch("/patch", f)

	if !called {
		t.Errorf("Expected f to be called, but it wasn't")
	}
}
