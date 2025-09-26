package web

import (
	"fmt"
	"net/http"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestCtrlGet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}
	ctrl.CtrlGet("/test", func(ctx *beecontext.Context) {
		// test logic
	})

	// assertions
	if _, ok := ctrl.routers[http.MethodGet]; !ok {
		t.Error("Expected method GET to be added to routers")
	}
}
