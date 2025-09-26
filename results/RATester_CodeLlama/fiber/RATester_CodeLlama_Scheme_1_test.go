package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestScheme_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.fasthttp.Request.Header.Set(HeaderXForwardedProto, "https")
	c.fasthttp.Request.Header.Set(HeaderXForwardedSsl, "on")
	c.fasthttp.Request.Header.Set(HeaderXUrlScheme, "http")
	assert.Equal(t, schemeHTTPS, c.Scheme())
}
