package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAppend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	ctx.Append("Header", "value1", "value2")

	expected := "value1, value2"
	actual := string(ctx.fasthttp.Response.Header.Peek("Header"))

	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
