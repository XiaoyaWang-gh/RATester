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

	expected := &fasthttp.Response{}
	ctx.fasthttp.Response = *expected

	result := ctx.Response()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
