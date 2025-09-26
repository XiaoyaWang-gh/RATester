package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestProtocol_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{
			Request: fasthttp.Request{
				Header: fasthttp.RequestHeader{},
			},
		},
	}

	ctx.fasthttp.Request.Header.Set("Protocol", "HTTP/1.1")

	protocol := ctx.Protocol()

	if protocol != "HTTP/1.1" {
		t.Errorf("Expected Protocol to be 'HTTP/1.1', but got '%s'", protocol)
	}
}
