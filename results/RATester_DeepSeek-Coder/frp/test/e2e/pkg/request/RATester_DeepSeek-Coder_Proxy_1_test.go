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

	r := &Request{}
	url := "http://proxy.example.com:8080"
	r.Proxy(url)
	if r.proxyURL != url {
		t.Errorf("Expected proxyURL to be %s, but got %s", url, r.proxyURL)
	}
}
