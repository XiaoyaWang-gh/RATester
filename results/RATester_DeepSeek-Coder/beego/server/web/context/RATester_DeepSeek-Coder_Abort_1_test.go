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
			Status: 200,
		},
	}

	ctx.Abort(404, "Not Found")

	if ctx.Output.Status != 404 {
		t.Errorf("Expected status to be 404, got %d", ctx.Output.Status)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic, but didn't get one")
		} else if r != "Not Found" {
			t.Errorf("Expected panic message to be 'Not Found', got '%v'", r)
		}
	}()

	ctx.Abort(404, "Not Found")
}
