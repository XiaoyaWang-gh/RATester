package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}
	assert.Equal(t, &c.fasthttp.Response, c.Response())
}
