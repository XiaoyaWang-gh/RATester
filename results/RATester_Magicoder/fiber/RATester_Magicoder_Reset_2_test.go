package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestReset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	ctx := &DefaultCtx{
		app: app,
	}

	fctx := &fasthttp.RequestCtx{}

	ctx.Reset(fctx)

	if ctx.indexRoute != -1 {
		t.Errorf("Expected indexRoute to be -1, but got %d", ctx.indexRoute)
	}

	if ctx.indexHandler != 0 {
		t.Errorf("Expected indexHandler to be 0, but got %d", ctx.indexHandler)
	}

	if ctx.matched != false {
		t.Errorf("Expected matched to be false, but got %v", ctx.matched)
	}

	if ctx.pathOriginal != "" {
		t.Errorf("Expected pathOriginal to be empty, but got %s", ctx.pathOriginal)
	}

	if ctx.method != "" {
		t.Errorf("Expected method to be empty, but got %s", ctx.method)
	}

	if ctx.methodINT != 0 {
		t.Errorf("Expected methodINT to be 0, but got %d", ctx.methodINT)
	}

	if ctx.fasthttp != fctx {
		t.Errorf("Expected fasthttp to be %v, but got %v", fctx, ctx.fasthttp)
	}

	if ctx.baseURI != "" {
		t.Errorf("Expected baseURI to be empty, but got %s", ctx.baseURI)
	}
}
