package fiber

import (
	"fmt"
	"testing"
)

func TestRestartRouting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	ctx := app.AcquireCtx(nil)

	err := ctx.RestartRouting()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	app.ReleaseCtx(ctx)
}
