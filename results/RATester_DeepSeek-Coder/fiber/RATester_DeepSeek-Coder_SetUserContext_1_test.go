package fiber

import (
	"context"
	"fmt"
	"testing"
)

func TestSetUserContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	app := New()
	ctx := app.AcquireCtx(nil)
	defer app.ReleaseCtx(ctx)

	testCtx := context.Background()
	ctx.SetUserContext(testCtx)

	if ctx.UserContext() != testCtx {
		t.Errorf("Expected %v, got %v", testCtx, ctx.UserContext())
	}
}
