package client

import (
	"fmt"
	"testing"
)

func TestC_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := C()

	// Test that the client is not nil
	if client == nil {
		t.Error("Expected client to not be nil")
	}

	// Test that the client has the expected default values
	if client.baseURL != "" {
		t.Errorf("Expected baseURL to be '', but got %s", client.baseURL)
	}

	// Add more tests as needed
}
