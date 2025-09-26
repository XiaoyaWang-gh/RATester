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

	r := &Response{}
	r.client = &Client{}
	r.request = &Request{}
	r.RawResponse = &fasthttp.Response{}
	r.cookie = []*fasthttp.Cookie{}

	r.Reset()

	if r.client != nil {
		t.Errorf("r.client should be nil")
	}

	if r.request != nil {
		t.Errorf("r.request should be nil")
	}

	if len(r.cookie) != 0 {
		t.Errorf("r.cookie should be empty")
	}

	if r.RawResponse != nil {
		t.Errorf("r.RawResponse should be nil")
	}
}
