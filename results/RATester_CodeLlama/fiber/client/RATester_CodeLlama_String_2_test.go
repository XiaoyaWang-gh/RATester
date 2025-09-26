package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	r.RawResponse = &fasthttp.Response{}
	r.RawResponse.SetBodyString("  hello world  ")
	if r.String() != "hello world" {
		t.Errorf("Response.String() = %v, want %v", r.String(), "hello world")
	}
}
