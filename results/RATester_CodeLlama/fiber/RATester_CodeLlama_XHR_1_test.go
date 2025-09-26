package fiber

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestXHR_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{}
	c.fasthttp = &fasthttp.RequestCtx{}
	c.fasthttp.Request.Header.Set(HeaderXRequestedWith, "xmlhttprequest")
	if !c.XHR() {
		t.Error("XHR() should return true")
	}
}
