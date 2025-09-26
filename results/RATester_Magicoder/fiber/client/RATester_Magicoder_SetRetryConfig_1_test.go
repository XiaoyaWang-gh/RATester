package client

import (
	"fmt"
	"testing"
)

func TestSetRetryConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	config := &RetryConfig{}

	client.SetRetryConfig(config)

	if client.retryConfig != config {
		t.Errorf("Expected client.retryConfig to be %v, but got %v", config, client.retryConfig)
	}
}
