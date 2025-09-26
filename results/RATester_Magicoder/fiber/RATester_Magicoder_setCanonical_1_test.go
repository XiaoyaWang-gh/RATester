package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestsetCanonical_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	key := "testKey"
	val := "testVal"

	ctx.setCanonical(key, val)

	header := ctx.fasthttp.Response.Header.Peek(key)
	if string(header) != val {
		t.Errorf("Expected %s, got %s", val, string(header))
	}
}
