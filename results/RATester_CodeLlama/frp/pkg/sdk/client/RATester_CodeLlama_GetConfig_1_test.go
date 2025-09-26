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
		address: "127.0.0.1:8080",
	}
	config, err := client.GetConfig(ctx)
	if err != nil {
		t.Errorf("GetConfig error: %v", err)
	}
	t.Logf("GetConfig: %s", config)
}
