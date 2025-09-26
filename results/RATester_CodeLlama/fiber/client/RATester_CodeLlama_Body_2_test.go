package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestBody_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	r.RawResponse = &fasthttp.Response{}
	r.RawResponse.SetBody([]byte("hello"))
	if string(r.Body()) != "hello" {
		t.Errorf("Body() = %v, want %v", string(r.Body()), "hello")
	}
}
