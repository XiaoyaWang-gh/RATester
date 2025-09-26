package request

import (
	"fmt"
	"testing"
)

func TestNew_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := New()

	if req.protocol != "tcp" {
		t.Errorf("Expected protocol to be 'tcp', but got '%s'", req.protocol)
	}

	if req.addr != "127.0.0.1" {
		t.Errorf("Expected addr to be '127.0.0.1', but got '%s'", req.addr)
	}

	if req.method != "GET" {
		t.Errorf("Expected method to be 'GET', but got '%s'", req.method)
	}

	if req.path != "/" {
		t.Errorf("Expected path to be '/', but got '%s'", req.path)
	}

	if len(req.headers) != 0 {
		t.Errorf("Expected headers to be empty, but got '%v'", req.headers)
	}
}
