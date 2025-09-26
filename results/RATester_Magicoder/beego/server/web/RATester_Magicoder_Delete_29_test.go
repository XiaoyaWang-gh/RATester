package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestDelete_29(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	rootpath := "/delete"
	f := func(ctx *beecontext.Context) {
		// test logic
	}

	n.Delete(rootpath, f)

	// assert
	if _, ok := n.handlers.routers[rootpath]; !ok {
		t.Errorf("Expected %s to be deleted", rootpath)
	}
}
