package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestIP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.app = &App{}
	c.app.config.ProxyHeader = "X-Forwarded-For"
	c.fasthttp.Request.Header.Set("X-Forwarded-For", "127.0.0.1")
	assert.Equal(t, "127.0.0.1", c.IP())
}
