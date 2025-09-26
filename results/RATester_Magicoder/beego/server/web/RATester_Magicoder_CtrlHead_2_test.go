package web

import (
	"fmt"
	"net/http"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestCtrlHead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}
	ctrl.CtrlHead("/test", func(ctx *beecontext.Context) {
		// test logic
	})

	// assertions
	if _, ok := ctrl.routers[http.MethodHead]; !ok {
		t.Error("Expected method HEAD to be added to routers")
	}
}
