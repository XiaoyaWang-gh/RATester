package fiber

import (
	"context"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestUserContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	expectedCtx := context.Background()
	ctx.SetUserContext(expectedCtx)

	actualCtx := ctx.UserContext()
	if actualCtx != expectedCtx {
		t.Errorf("Expected context %v, but got %v", expectedCtx, actualCtx)
	}
}
