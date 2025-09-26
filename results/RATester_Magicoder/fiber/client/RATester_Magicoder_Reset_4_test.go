package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestReset_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{
		client:      &Client{},
		request:     &Request{},
		RawResponse: &fasthttp.Response{},
		cookie: []*fasthttp.Cookie{
			{},
		},
	}

	r.Reset()

	if r.client != nil {
		t.Error("Expected client to be nil after reset")
	}
	if r.request != nil {
		t.Error("Expected request to be nil after reset")
	}
	if r.RawResponse.Header.Len() != 0 {
		t.Error("Expected RawResponse to be empty after reset")
	}
	if len(r.cookie) != 0 {
		t.Error("Expected cookie to be empty after reset")
	}
}
