package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		fasthttp: &fasthttp.RequestCtx{},
	}

	expected := &fasthttp.Request{}
	actual := ctx.Request()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
