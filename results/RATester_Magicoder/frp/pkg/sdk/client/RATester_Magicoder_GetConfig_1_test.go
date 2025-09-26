package client

import (
	"context"
	"fmt"
	"testing"
)

func TestGetConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	client := &Client{
		address: "localhost:8080",
	}
	config, err := client.GetConfig(ctx)
	if err != nil {
		t.Fatalf("GetConfig failed: %v", err)
	}
	if config == "" {
		t.Fatalf("GetConfig returned empty config")
	}
}
