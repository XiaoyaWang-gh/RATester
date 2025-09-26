package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestFresh_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.fasthttp.Request.Header.Set(HeaderIfModifiedSince, "")
	ctx.fasthttp.Request.Header.Set(HeaderIfNoneMatch, "")

	result := ctx.Fresh()

	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
