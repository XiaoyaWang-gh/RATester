package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestHostname_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.fasthttp.Request.SetRequestURI("http://example.com")

	hostname := ctx.Hostname()

	if hostname != "example.com" {
		t.Errorf("Expected hostname to be 'example.com', but got '%s'", hostname)
	}
}
