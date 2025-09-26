package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestHeader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{
		RawResponse: &fasthttp.Response{},
	}

	r.RawResponse.Header.Set("Content-Type", "application/json")

	expected := "application/json"
	actual := r.Header("Content-Type")

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
