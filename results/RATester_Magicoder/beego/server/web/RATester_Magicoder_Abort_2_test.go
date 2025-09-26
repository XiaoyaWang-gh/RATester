package web

import (
	"fmt"
	"testing"
)

func TestAbort_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{}
	ctrl.Abort("200")
	if ctrl.Ctx.Output.Status != 200 {
		t.Errorf("Expected status 200, got %d", ctrl.Ctx.Output.Status)
	}

	ctrl.Abort("not a number")
	if ctrl.Ctx.Output.Status != 200 {
		t.Errorf("Expected status 200, got %d", ctrl.Ctx.Output.Status)
	}
}
