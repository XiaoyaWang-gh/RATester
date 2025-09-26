package client

import (
	"context"
	"fmt"
	"testing"
)

func TestUpdateConfig_1(t *testing.T) {
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

	content := `{"key": "value"}`
	err := client.UpdateConfig(ctx, content)
	if err != nil {
		t.Errorf("UpdateConfig failed: %v", err)
	}
}
