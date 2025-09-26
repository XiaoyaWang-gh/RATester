package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestURLFor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrl.AddRouterMethod("GET", "/test", func(ctx *beecontext.Context) {})

	url := ctrl.URLFor("test.controller.method", "key", "value")
	if url != "/test" {
		t.Errorf("Expected /test, got %s", url)
	}

	url = ctrl.URLFor("test.controller", "key", "value")
	if url != "" {
		t.Errorf("Expected empty string, got %s", url)
	}

	url = ctrl.URLFor("test.controller.method", "key")
	if url != "" {
		t.Errorf("Expected empty string, got %s", url)
	}
}
