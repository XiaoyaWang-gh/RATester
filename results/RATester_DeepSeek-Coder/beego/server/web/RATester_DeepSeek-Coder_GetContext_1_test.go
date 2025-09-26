package web

import (
	"fmt"
	"sync"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestGetContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		pool: sync.Pool{
			New: func() interface{} {
				return &beecontext.Context{}
			},
		},
	}

	ctx := ctrl.GetContext()

	if ctx == nil {
		t.Errorf("Expected non-nil context, got nil")
	}

	if ctx.Input == nil || ctx.Output == nil || ctx.Request == nil || ctx.ResponseWriter == nil {
		t.Errorf("Expected initialized context fields, got nil")
	}
}
