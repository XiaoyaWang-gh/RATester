package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	resp := ctx.Response()
	if resp == nil {
		t.Errorf("Expected non-nil response, got nil")
	}

	if resp != &ctx.fasthttp.Response {
		t.Errorf("Expected response to be equal to &ctx.fasthttp.Response, got %v", resp)
	}
}
