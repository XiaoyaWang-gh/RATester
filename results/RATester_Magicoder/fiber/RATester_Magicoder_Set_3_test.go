package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSet_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	key := "key"
	val := "value"

	ctx.Set(key, val)

	header := ctx.fasthttp.Response.Header.Peek(key)
	if string(header) != val {
		t.Errorf("Expected %s, got %s", val, string(header))
	}
}
