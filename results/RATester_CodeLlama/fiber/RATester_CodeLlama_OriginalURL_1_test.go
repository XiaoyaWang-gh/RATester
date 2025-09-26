package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestOriginalURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.fasthttp.Request.Header.SetRequestURI("http://localhost:8080/test")
	assert.Equal(t, "http://localhost:8080/test", c.OriginalURL())
}
