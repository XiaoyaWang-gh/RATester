package visitor

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestOpenConn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	qs := &QUICTunnelSession{
		listenConn: &net.UDPConn{},
	}
	// CONTEXT_1
	ctx := context.Background()
	// CONTEXT_2
	conn, err := qs.OpenConn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	// CONTEXT_3
	if qs.listenConn != conn {
		t.Fatal("listenConn not equal conn")
	}
	// CONTEXT_4
	if err != nil {
		t.Fatal(err)
	}
}
