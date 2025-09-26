package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestCtrlOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}
	ctrl.CtrlOptions("/test", func(ctx *beecontext.Context) {
		// test logic here
	})

	// assertions here
}
