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
			URL:    &url.URL{Path: "/path"},
			Header: make(http.Header),
		},
	}

	result := req.GetRequest()

	if result.Method != "GET" {
		t.Errorf("Expected method to be 'GET', but got '%s'", result.Method)
	}

	if result.URL.String() != "http://example.com/path" {
		t.Errorf("Expected URL to be 'http://example.com/path', but got '%s'", result.URL.String())
	}

	if result.Header.Get("Content-Type") != "" {
		t.Errorf("Expected header 'Content-Type' to be empty, but got '%s'", result.Header.Get("Content-Type"))
	}
}
