package httplib

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetProtocolVersion_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &BeegoHTTPRequest{
		url: "http://example.com",
		req: &http.Request{},
	}

	// Test with a valid version
	req.SetProtocolVersion("HTTP/1.0")
	if req.req.Proto != "HTTP/1.0" {
		t.Errorf("Expected Proto to be 'HTTP/1.0', but got '%s'", req.req.Proto)
	}

	// Test with an empty version
	req.SetProtocolVersion("")
	if req.req.Proto != "HTTP/1.1" {
		t.Errorf("Expected Proto to be 'HTTP/1.1', but got '%s'", req.req.Proto)
	}

	// Test with an invalid version
	req.SetProtocolVersion("HTTP/3.0")
	if req.req.Proto != "HTTP/1.1" {
		t.Errorf("Expected Proto to be 'HTTP/1.1', but got '%s'", req.req.Proto)
	}
}
