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
		// Test logic here
	})

	if _, ok := ctrl.routers[http.MethodPost]; !ok {
		t.Error("Expected method post to be added")
	}
}
