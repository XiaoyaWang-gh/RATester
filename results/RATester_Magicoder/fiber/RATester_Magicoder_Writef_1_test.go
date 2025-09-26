package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestWritef_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{}
	ctx.fasthttp = &fasthttp.RequestCtx{}
	ctx.fasthttp.Response.BodyWriter()

	_, err := ctx.Writef("Hello, %s!", "world")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "Hello, world!"
	if string(ctx.fasthttp.Response.Body()) != expected {
		t.Errorf("Expected %q, got %q", expected, ctx.fasthttp.Response.Body())
	}
}
