package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestPatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}
	ctrl.Patch("/test", func(ctx *beecontext.Context) {
		// test logic
	})

	// assert
	if _, ok := ctrl.routers["patch"]; !ok {
		t.Error("Patch method not added to routers")
	}
}
