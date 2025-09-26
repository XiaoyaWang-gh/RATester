package request

import (
	"fmt"
	"reflect"
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
	headers := map[string]string{"Content-Type": "application/json"}
	r.HTTPParams(method, host, path, headers)

	if r.method != method {
		t.Errorf("Expected method %s, but got %s", method, r.method)
	}
	if r.host != host {
		t.Errorf("Expected host %s, but got %s", host, r.host)
	}
	if r.path != path {
		t.Errorf("Expected path %s, but got %s", path, r.path)
	}
	if !reflect.DeepEqual(r.headers, headers) {
		t.Errorf("Expected headers %v, but got %v", headers, r.headers)
	}
}
