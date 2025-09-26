package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestCookies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := New()
	req := fasthttp.AcquireRequest()
	req.Header.SetCookie("test", "value")
	fctx := fasthttp.RequestCtx{Request: *req}
	ctx := &DefaultCtx{
		app:      app,
		fasthttp: &fctx,
	}

	result := ctx.Cookies("test", "default")
	if result != "value" {
		t.Errorf("Expected 'value', got '%s'", result)
	}

	result = ctx.Cookies("not_existing", "default")
	if result != "default" {
		t.Errorf("Expected 'default', got '%s'", result)
	}
}
