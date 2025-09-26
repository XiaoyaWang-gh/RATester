package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestIPs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.fasthttp.Request.Header.Set(HeaderXForwardedFor, "192.168.1.1, 192.168.1.2")
	assert.Equal(t, []string{"192.168.1.1", "192.168.1.2"}, c.IPs())
}
