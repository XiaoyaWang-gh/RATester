package context

import (
	"fmt"
	"testing"
)

func TestRender_2(t *testing.T) {
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

	statusCode := StatusCode(200)
	statusCode.Render(ctx)

	if ctx.Output.Status != 200 {
		t.Errorf("Expected status code to be 200, but got %d", ctx.Output.Status)
	}
}
