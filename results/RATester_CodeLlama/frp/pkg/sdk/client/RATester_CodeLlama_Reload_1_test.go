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
	c := &Client{}
	err := c.Reload(ctx, true)
	if err != nil {
		t.Errorf("Reload failed, err: %v", err)
	}
}
