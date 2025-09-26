package client

import (
	"fmt"
	"testing"
)

func TestGetBadResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	res := getBadResponse()

	if res.StatusCode != 407 {
		t.Errorf("Expected status code 407, but got %d", res.StatusCode)
	}

	header := res.Header
	proxyAuth, ok := header["Proxy-Authenticate"]
	if !ok || proxyAuth[0] != "Basic" {
		t.Errorf("Expected Proxy-Authenticate header to be 'Basic', but got %s", proxyAuth[0])
	}

	connection, ok := header["Connection"]
	if !ok || connection[0] != "close" {
		t.Errorf("Expected Connection header to be 'close', but got %s", connection[0])
	}
}
