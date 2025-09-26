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
	headers := map[string]string{"Content-Type": "application/json"}
	r.HTTPHeaders(headers)

	if r.headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type to be application/json, but got %s", r.headers["Content-Type"])
	}
}
