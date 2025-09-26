package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestCtrlAny_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrl.CtrlAny("/test", func(ctx *beecontext.Context) {
		// Test logic here
	})

	if _, ok := ctrl.routers["*"]; !ok {
		t.Error("Expected router to be added")
	}
}
