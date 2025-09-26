package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestClearCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.fasthttp.Request.Header.SetCookie("test", "value")
	ctx.fasthttp.Request.Header.SetCookie("test2", "value2")

	ctx.ClearCookie("test")

	if ctx.fasthttp.Request.Header.Cookie("test") != nil {
		t.Error("Expected cookie to be cleared")
	}

	if ctx.fasthttp.Request.Header.Cookie("test2") == nil {
		t.Error("Expected cookie to be preserved")
	}
}
