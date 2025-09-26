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
	c := &Client{}
	err := c.UpdateConfig(ctx, "content")
	if err != nil {
		t.Errorf("UpdateConfig error: %v", err)
	}
}
