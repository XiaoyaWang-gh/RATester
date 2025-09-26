package client

import (
	"context"
	"fmt"
	"testing"
)

func TestStop_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	client := &Client{
		address: "localhost:8080",
	}

	err := client.Stop(ctx)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
