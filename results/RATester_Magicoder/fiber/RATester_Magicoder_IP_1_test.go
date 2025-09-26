package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestIP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	app.config.ProxyHeader = "X-Forwarded-For"
	ctx := &DefaultCtx{
		app:      app,
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.fasthttp.Request.Header.Set("X-Forwarded-For", "192.168.1.1")

	ip := ctx.IP()

	if ip != "192.168.1.1" {
		t.Errorf("Expected IP to be 192.168.1.1, but got %s", ip)
	}
}
