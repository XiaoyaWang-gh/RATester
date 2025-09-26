package fiber

import (
	"fmt"
	"testing"
)

func TestApp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		app: &App{},
	}

	app := ctx.App()

	if app != ctx.app {
		t.Errorf("Expected %v, got %v", ctx.app, app)
	}
}
