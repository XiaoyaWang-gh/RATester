package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestCookies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.fasthttp.Request.Header.SetCookie("key", "value")
	assert.Equal(t, "value", c.Cookies("key"))
}
