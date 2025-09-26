package context

import (
	"fmt"
	"testing"
)

func TestReset_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			Context: &Context{},
		},
		Output: &BeegoOutput{},
	}

	ctx.Output.Reset(ctx)

	if ctx.Output.Context != ctx {
		t.Errorf("Expected ctx.Output.Context to be %v, got %v", ctx, ctx.Output.Context)
	}

	if ctx.Output.Status != 0 {
		t.Errorf("Expected ctx.Output.Status to be 0, got %v", ctx.Output.Status)
	}
}
