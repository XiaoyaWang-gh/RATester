package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Header{
		RequestHeader: &fasthttp.RequestHeader{},
	}

	r := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	h.SetHeaders(r)

	if len(h.RequestHeader.Peek("key1")) == 0 || string(h.RequestHeader.Peek("key1")) != "value1" {
		t.Errorf("Expected key1 to be set to value1")
	}

	if len(h.RequestHeader.Peek("key2")) == 0 || string(h.RequestHeader.Peek("key2")) != "value2" {
		t.Errorf("Expected key2 to be set to value2")
	}
}
