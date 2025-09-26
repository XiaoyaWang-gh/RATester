package client

import (
	"context"
	"fmt"
	"testing"
)

func TestSetConfigToRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &Request{}
	config := []Config{
		{
			Ctx: context.Background(),
		},
	}
	setConfigToRequest(req, config...)
	if req.ctx != context.Background() {
		t.Errorf("req.ctx = %v, want %v", req.ctx, context.Background())
	}
}
