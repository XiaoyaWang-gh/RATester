package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rootpath := "/test"
	f := func(ctx *beecontext.Context) {
		// test function implementation
	}

	BeeApp := Options(rootpath, f)

	if BeeApp.Handlers == nil {
		t.Error("Expected Handlers to be not nil")
	}

	if BeeApp.Server == nil {
		t.Error("Expected Server to be not nil")
	}

	if BeeApp.Cfg == nil {
		t.Error("Expected Cfg to be not nil")
	}

	if len(BeeApp.LifeCycleCallbacks) != 0 {
		t.Error("Expected LifeCycleCallbacks to be empty")
	}
}
