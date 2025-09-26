package fiber

import (
	"fmt"
	"sync"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestReleaseCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		pool: sync.Pool{
			New: func() any {
				return &DefaultCtx{}
			},
		},
	}

	ctx := app.pool.Get().(*DefaultCtx)
	ctx.fasthttp = &fasthttp.RequestCtx{}

	app.ReleaseCtx(ctx)

	if ctx.fasthttp != nil {
		t.Error("fasthttp should be nil after release")
	}

	if app.pool.New() != nil {
		t.Error("pool should be empty after release")
	}
}
