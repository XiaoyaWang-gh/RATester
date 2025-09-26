package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAddHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Header{
		RequestHeader: &fasthttp.RequestHeader{},
	}

	r := map[string][]string{
		"key1": {"value1", "value2"},
		"key2": {"value3", "value4"},
	}

	h.AddHeaders(r)

	if len(h.RequestHeader.Peek("key1")) == 0 || len(h.RequestHeader.Peek("key2")) == 0 {
		t.Error("AddHeaders failed")
	}
}
