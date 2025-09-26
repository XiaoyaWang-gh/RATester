package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestAbort_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	ctrl.Abort("200")
	if ctrl.Ctx.ResponseWriter.Status != 200 {
		t.Errorf("Expected status code 200, got %d", ctrl.Ctx.ResponseWriter.Status)
	}

	ctrl.Abort("500")
	if ctrl.Ctx.ResponseWriter.Status != 500 {
		t.Errorf("Expected status code 500, got %d", ctrl.Ctx.ResponseWriter.Status)
	}

	ctrl.Abort("invalid")
	if ctrl.Ctx.ResponseWriter.Status != 200 {
		t.Errorf("Expected status code 200, got %d", ctrl.Ctx.ResponseWriter.Status)
	}
}
