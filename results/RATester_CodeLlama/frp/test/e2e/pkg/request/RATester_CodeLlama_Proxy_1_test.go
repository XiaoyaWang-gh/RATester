package request

import (
	"fmt"
	"testing"
)

func TestProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		proxyURL: "http://127.0.0.1:8080",
	}
	r.Proxy("http://127.0.0.1:8080")
	if r.proxyURL != "http://127.0.0.1:8080" {
		t.Errorf("Proxy() failed")
	}
}
