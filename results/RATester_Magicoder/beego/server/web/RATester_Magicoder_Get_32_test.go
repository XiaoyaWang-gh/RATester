package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestGet_32(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	rootpath := "/get"
	f := func(ctx *beecontext.Context) {
		// test logic
	}

	n.Get(rootpath, f)

	// assert
	if _, ok := n.handlers.routers["GET"]; !ok {
		t.Error("Expected GET router to be added")
	}
}
