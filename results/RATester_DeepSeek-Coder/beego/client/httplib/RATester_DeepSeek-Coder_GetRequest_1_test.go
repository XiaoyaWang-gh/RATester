package httplib

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestGetRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		url: "http://example.com",
		req: &http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme: "http",
				Host:   "example.com",
			},
		},
	}

	httpReq := req.GetRequest()

	if httpReq.Method != "GET" {
		t.Errorf("Expected method to be 'GET', got '%s'", httpReq.Method)
	}

	if httpReq.URL.String() != "http://example.com" {
		t.Errorf("Expected URL to be 'http://example.com', got '%s'", httpReq.URL.String())
	}
}
