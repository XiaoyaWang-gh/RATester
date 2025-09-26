package request

import (
	"fmt"
	"testing"
)

func TestHTTPParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	method := "GET"
	host := "example.com"
	path := "/path"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	r.HTTPParams(method, host, path, headers)

	if r.method != method {
		t.Errorf("Expected method %s, got %s", method, r.method)
	}
	if r.host != host {
		t.Errorf("Expected host %s, got %s", host, r.host)
	}
	if r.path != path {
		t.Errorf("Expected path %s, got %s", path, r.path)
	}
	if r.headers["Content-Type"] != "application/json" {
		t.Errorf("Expected header Content-Type %s, got %s", "application/json", r.headers["Content-Type"])
	}
}
