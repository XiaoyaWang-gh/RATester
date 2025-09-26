package context

import (
	"fmt"
	"testing"
)

func TestAbort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Output: &BeegoOutput{
			Context: &Context{},
		},
	}

	ctx.Abort(404, "Not Found")

	if ctx.Output.Status != 404 {
		t.Errorf("Expected status 404, got %d", ctx.Output.Status)
	}
}
