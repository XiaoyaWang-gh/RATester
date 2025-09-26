package pmux

import (
	"fmt"
	"net"
	"testing"
)

func TestLocalAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	conn, err := net.Dial(listener.Addr().Network(), listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	pConn := &PortConn{Conn: conn}

	addr := pConn.LocalAddr()
	if addr == nil {
		t.Fatal("LocalAddr() returned nil")
	}

	if addr.Network() != "tcp" {
		t.Errorf("Network() = %s; want tcp", addr.Network())
	}

	if addr.String() == "" {
		t.Error("String() returned an empty string")
	}
}
