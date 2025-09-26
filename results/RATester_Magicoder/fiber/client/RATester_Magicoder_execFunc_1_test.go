package client

import (
	"context"
	"fmt"
	"testing"
)

func TestexecFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	req := &Request{}
	ctx := context.Background()

	core := &core{
		client: client,
		req:    req,
		ctx:    ctx,
	}

	resp, err := core.execFunc()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp == nil {
		t.Error("Expected non-nil response, got nil")
	}
}
