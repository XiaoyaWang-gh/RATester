package client

import (
	"fmt"
	"testing"
)

func TestClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		client: &Client{},
	}

	if r.Client() != r.client {
		t.Errorf("Expected Client() to return %v, but got %v", r.client, r.Client())
	}
}
