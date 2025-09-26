package http

import (
	"fmt"
	"testing"
)

func TestProxyUnauthorizedResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	res := ProxyUnauthorizedResponse()

	if res.StatusCode != 407 {
		t.Errorf("Expected status code 407, but got %d", res.StatusCode)
	}

	if res.Header.Get("Proxy-Authenticate") != `Basic realm="Restricted"` {
		t.Errorf("Expected Proxy-Authenticate header to be 'Basic realm=\"Restricted\"', but got '%s'", res.Header.Get("Proxy-Authenticate"))
	}
}
