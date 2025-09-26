package web

import (
	"fmt"
	"net/http"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestCtrlDelete_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrl.CtrlDelete("/test", func(ctx *beecontext.Context) {
		// Test logic here
	})

	if _, ok := ctrl.routers[http.MethodDelete]; !ok {
		t.Error("Expected method delete to be added to routers")
	}
}
