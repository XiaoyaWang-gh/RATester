package client

import (
	"context"
	"fmt"
	"testing"
)

func TestReload_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	client := &Client{
		address:  "localhost:8080",
		authUser: "user",
		authPwd:  "password",
	}

	// Test with strictMode = true
	err := client.Reload(ctx, true)
	if err != nil {
		t.Errorf("Reload with strictMode = true failed: %v", err)
	}

	// Test with strictMode = false
	err = client.Reload(ctx, false)
	if err != nil {
		t.Errorf("Reload with strictMode = false failed: %v", err)
	}
}
