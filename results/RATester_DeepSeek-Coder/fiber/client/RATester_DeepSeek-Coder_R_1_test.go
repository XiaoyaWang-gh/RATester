package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestR_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	client := &Client{
		fasthttp: &fasthttp.Client{},
	}

	req := client.R()

	if req.client != client {
		t.Errorf("Expected client to be %v, got %v", client, req.client)
	}

	if req.header == nil {
		t.Errorf("Expected header to be not nil")
	}

	if req.params == nil {
		t.Errorf("Expected params to be not nil")
	}

	if req.cookies == nil {
		t.Errorf("Expected cookies to be not nil")
	}

	if req.path == nil {
		t.Errorf("Expected path to be not nil")
	}

	if req.url != "" {
		t.Errorf("Expected url to be empty, got %v", req.url)
	}

	if req.method != "" {
		t.Errorf("Expected method to be empty, got %v", req.method)
	}
}
