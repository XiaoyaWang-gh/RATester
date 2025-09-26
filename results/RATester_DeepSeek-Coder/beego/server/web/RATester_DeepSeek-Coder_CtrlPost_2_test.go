package web

import (
	"fmt"
	"net/http"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestCtrlPost_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrl.CtrlPost("/test", func(ctx *beecontext.Context) {
		ctx.Output.SetStatus(200)
	})

	_, ok := ctrl.routers[http.MethodPost]
	if !ok {
		t.Errorf("Expected method %s to be registered", http.MethodPost)
	}
}
