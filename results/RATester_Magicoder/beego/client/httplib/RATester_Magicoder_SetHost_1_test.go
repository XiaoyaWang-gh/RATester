package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		url: "http://example.com",
		req: &http.Request{},
	}

	host := "test.com"
	req.SetHost(host)

	if req.req.Host != host {
		t.Errorf("Expected host to be %s, but got %s", host, req.req.Host)
	}
}
