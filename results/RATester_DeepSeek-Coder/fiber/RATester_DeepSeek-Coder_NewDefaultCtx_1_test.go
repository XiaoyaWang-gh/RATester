package fiber

import (
	"fmt"
	"testing"
)

func TestNewDefaultCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	ctx := NewDefaultCtx(app)

	if ctx.app != app {
		t.Errorf("Expected app to be %v, got %v", app, ctx.app)
	}
}
