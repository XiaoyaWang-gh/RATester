package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestReleaseResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	resp := &Response{
		RawResponse: &fasthttp.Response{},
	}

	ReleaseResponse(resp)

	if resp.RawResponse != nil {
		t.Errorf("Expected RawResponse to be nil after ReleaseResponse, but got %v", resp.RawResponse)
	}
}
