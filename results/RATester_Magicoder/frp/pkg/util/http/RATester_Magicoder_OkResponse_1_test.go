package http

import (
	"fmt"
	"testing"
)

func TestOkResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	res := OkResponse()

	if res.Status != "OK" {
		t.Errorf("Expected status 'OK', got '%s'", res.Status)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	if res.Proto != "HTTP/1.1" {
		t.Errorf("Expected protocol 'HTTP/1.1', got '%s'", res.Proto)
	}

	if res.ProtoMajor != 1 {
		t.Errorf("Expected ProtoMajor 1, got %d", res.ProtoMajor)
	}

	if res.ProtoMinor != 1 {
		t.Errorf("Expected ProtoMinor 1, got %d", res.ProtoMinor)
	}
}
