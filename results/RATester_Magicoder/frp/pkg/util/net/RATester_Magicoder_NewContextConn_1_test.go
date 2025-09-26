package net

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestNewContextConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	cc := NewContextConn(ctx, conn)

	if cc.ctx != ctx {
		t.Errorf("Expected context to be %v, but got %v", ctx, cc.ctx)
	}

	if cc.Conn != conn {
		t.Errorf("Expected connection to be %v, but got %v", conn, cc.Conn)
	}
}
