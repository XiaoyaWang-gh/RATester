package fiber

import (
	"fmt"
	"net"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.fasthttp.RemoteAddr().(*net.TCPAddr).Port = 8080
	assert.Equal(t, "8080", c.Port())
}
