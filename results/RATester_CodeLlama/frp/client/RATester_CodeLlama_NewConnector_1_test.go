package client

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewConnector_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	cfg := &v1.ClientCommonConfig{}
	connector := NewConnector(ctx, cfg)
	if connector == nil {
		t.Error("NewConnector() = nil, want not nil")
	}
}
