package vhost

import (
	"fmt"
	"io"
	"testing"
)

func TestNotFoundResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	res := NotFoundResponse()

	if res.Status != "Not Found" {
		t.Errorf("Expected status 'Not Found', got '%s'", res.Status)
	}

	if res.StatusCode != 404 {
		t.Errorf("Expected status code 404, got %d", res.StatusCode)
	}

	if res.Proto != "HTTP/1.1" {
		t.Errorf("Expected protocol 'HTTP/1.1', got '%s'", res.Proto)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}

	if len(content) == 0 {
		t.Errorf("Expected non-empty body")
	}
}
