package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSetHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		header: &Header{
			RequestHeader: &fasthttp.RequestHeader{},
		},
	}

	key := "Content-Type"
	val := "application/json"

	r.SetHeader(key, val)

	if got := r.header.Peek(key); string(got) != val {
		t.Errorf("SetHeader() = %v, want %v", got, val)
	}
}
