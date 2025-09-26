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

	s := StatusCode(200)
	ctx := &Context{}
	s.Render(ctx)
	if ctx.Output.Status != 200 {
		t.Errorf("StatusCode.Render() = %v, want %v", ctx.Output.Status, 200)
	}
}
