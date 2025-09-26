package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestStatusCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	r.RawResponse = &fasthttp.Response{}
	r.RawResponse.SetStatusCode(200)
	if r.StatusCode() != 200 {
		t.Errorf("StatusCode() = %v, want %v", r.StatusCode(), 200)
	}
}
