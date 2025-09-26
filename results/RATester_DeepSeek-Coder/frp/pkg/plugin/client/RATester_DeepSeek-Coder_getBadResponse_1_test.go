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

	if res.Status != "407 Not authorized" {
		t.Errorf("Expected status '407 Not authorized', got %s", res.Status)
	}

	if res.StatusCode != 407 {
		t.Errorf("Expected status code 407, got %d", res.StatusCode)
	}

	if _, ok := res.Header["Proxy-Authenticate"]; !ok {
		t.Errorf("Expected header 'Proxy-Authenticate' not found")
	}

	if _, ok := res.Header["Connection"]; !ok {
		t.Errorf("Expected header 'Connection' not found")
	}
}
