package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAcceptsCharsets_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.fasthttp.Request.Header.Set(HeaderAcceptCharset, "utf-8")

	result := ctx.AcceptsCharsets("utf-8")

	if result != "utf-8" {
		t.Errorf("Expected 'utf-8', got '%s'", result)
	}
}
