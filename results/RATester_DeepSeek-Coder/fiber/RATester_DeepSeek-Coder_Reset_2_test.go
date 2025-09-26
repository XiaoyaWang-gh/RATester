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
	fctx := &fasthttp.RequestCtx{}
	c := &DefaultCtx{
		app:      app,
		fasthttp: fctx,
	}

	c.Reset(fctx)

	if c.indexRoute != -1 {
		t.Errorf("Expected indexRoute to be -1, got %d", c.indexRoute)
	}

	if c.indexHandler != 0 {
		t.Errorf("Expected indexHandler to be 0, got %d", c.indexHandler)
	}

	if c.matched != false {
		t.Errorf("Expected matched to be false, got %v", c.matched)
	}

	if c.app != app {
		t.Errorf("Expected app to be %v, got %v", app, c.app)
	}

	if c.fasthttp != fctx {
		t.Errorf("Expected fasthttp to be %v, got %v", fctx, c.fasthttp)
	}
}
