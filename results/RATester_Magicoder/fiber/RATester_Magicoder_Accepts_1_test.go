package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAccepts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.fasthttp.Request.Header.Set(HeaderAccept, "text/html")

	result := ctx.Accepts("text/html")

	if result != "text/html" {
		t.Errorf("Expected 'text/html', got '%s'", result)
	}
}
