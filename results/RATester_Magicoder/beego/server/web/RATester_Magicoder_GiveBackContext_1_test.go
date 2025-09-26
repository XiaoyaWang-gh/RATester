package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestGiveBackContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}
	ctx := &beecontext.Context{}
	ctrl.GiveBackContext(ctx)

	// Assert that the context was put back into the pool
	if ctrl.pool.Get() != ctx {
		t.Errorf("Expected context to be put back into the pool")
	}
}
