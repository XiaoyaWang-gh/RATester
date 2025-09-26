package context

import (
	"fmt"
	"testing"
)

func TestReset_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		Input: &BeegoInput{
			Context: &Context{
				Input: &BeegoInput{},
			},
		},
	}

	ctx.Input.Reset(ctx)

	if ctx.Input.Context != ctx {
		t.Errorf("Expected ctx.Input.Context to be %v, but got %v", ctx, ctx.Input.Context)
	}
	if ctx.Input.CruSession != nil {
		t.Errorf("Expected ctx.Input.CruSession to be nil, but got %v", ctx.Input.CruSession)
	}
	if len(ctx.Input.pnames) != 0 {
		t.Errorf("Expected ctx.Input.pnames to be empty, but got %v", ctx.Input.pnames)
	}
	if len(ctx.Input.pvalues) != 0 {
		t.Errorf("Expected ctx.Input.pvalues to be empty, but got %v", ctx.Input.pvalues)
	}
	if ctx.Input.data != nil {
		t.Errorf("Expected ctx.Input.data to be nil, but got %v", ctx.Input.data)
	}
	if len(ctx.Input.RequestBody) != 0 {
		t.Errorf("Expected ctx.Input.RequestBody to be empty, but got %v", ctx.Input.RequestBody)
	}
}
