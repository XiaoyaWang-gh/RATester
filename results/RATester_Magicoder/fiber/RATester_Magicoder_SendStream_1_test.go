package fiber

import (
	"fmt"
	"strings"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSendStream_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	reader := strings.NewReader("test")
	err := ctx.SendStream(reader)

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if string(ctx.fasthttp.Response.Body()) != "test" {
		t.Errorf("Expected 'test', got %s", string(ctx.fasthttp.Response.Body()))
	}
}
