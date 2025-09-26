package client

import (
	"fmt"
	"testing"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	request := &Request{}
	response := &Response{
		client:  client,
		request: request,
	}

	response.Close()

	if response.request != nil {
		t.Errorf("Expected request to be nil after Close, got %v", response.request)
	}

	if response.client != nil {
		t.Errorf("Expected client to be nil after Close, got %v", response.client)
	}
}
