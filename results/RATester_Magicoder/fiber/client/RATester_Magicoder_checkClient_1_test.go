package client

import (
	"fmt"
	"testing"
)

func TestcheckClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{
		client: nil,
	}
	req.checkClient()
	if req.client != defaultClient {
		t.Errorf("Expected client to be set to defaultClient")
	}
}
