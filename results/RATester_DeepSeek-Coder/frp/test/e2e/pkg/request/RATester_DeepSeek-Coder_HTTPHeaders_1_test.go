package request

import (
	"fmt"
	"testing"
)

func TestHTTPHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token",
	}
	r.HTTPHeaders(headers)
	if r.headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type to be 'application/json', got %s", r.headers["Content-Type"])
	}
	if r.headers["Authorization"] != "Bearer token" {
		t.Errorf("Expected Authorization to be 'Bearer token', got %s", r.headers["Authorization"])
	}
}
