package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		req: &http.Request{
			Header: make(http.Header),
		},
	}

	key := "Content-Type"
	value := "application/json"

	req.Header(key, value)

	if req.req.Header.Get(key) != value {
		t.Errorf("Expected %s, got %s", value, req.req.Header.Get(key))
	}
}
