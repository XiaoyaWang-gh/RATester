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
	cfg := &v1.ClientCommonConfig{
		ServerAddr: "localhost",
		ServerPort: 8080,
	}

	connector := NewConnector(ctx, cfg)

	if connector == nil {
		t.Error("Expected a non-nil connector, but got nil")
	}

	err := connector.Open()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	conn, err := connector.Connect()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if conn == nil {
		t.Error("Expected a non-nil connection, but got nil")
	}

	err = connector.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
