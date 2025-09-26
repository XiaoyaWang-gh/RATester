package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestStopRun_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		} else if r != ErrAbort {
			t.Errorf("Expected panic with error %v, but got %v", ErrAbort, r)
		}
	}()

	ctrl.StopRun()
}
