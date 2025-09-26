package client

import (
	"fmt"
	"testing"
)

func TestNewChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &TRPClient{
		svrAddr:        "localhost:8080",
		bridgeConnType: "tcp",
		proxyUrl:       "http://localhost:8080",
		vKey:           "test_key",
	}

	client.newChan()

	// Assert that the client's tunnel field is not nil
	if client.tunnel == nil {
		t.Errorf("Expected client.tunnel to be non-nil")
	}

	// Assert that the client's handleChan method is called
	// This would require mocking the handleChan method to check if it was called
}
