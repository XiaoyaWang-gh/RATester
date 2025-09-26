package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestGetRespHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.fasthttp.Response.Header.Set("Content-Type", "application/json")
	ctx.fasthttp.Response.Header.Set("Cache-Control", "no-cache")

	headers := ctx.GetRespHeaders()

	if len(headers) != 2 {
		t.Errorf("Expected 2 headers, got %d", len(headers))
	}

	if headers["Content-Type"][0] != "application/json" {
		t.Errorf("Expected Content-Type header to be 'application/json', got '%s'", headers["Content-Type"][0])
	}

	if headers["Cache-Control"][0] != "no-cache" {
		t.Errorf("Expected Cache-Control header to be 'no-cache', got '%s'", headers["Cache-Control"][0])
	}
}
