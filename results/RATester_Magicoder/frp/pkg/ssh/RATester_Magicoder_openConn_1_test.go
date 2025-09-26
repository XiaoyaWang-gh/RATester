package ssh

import (
	"fmt"
	"net"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestOpenConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &TunnelServer{
		underlyingConn: &net.TCPConn{},
		sshConn:        &ssh.ServerConn{},
	}

	addr := &tcpipForward{
		Host: "localhost",
		Port: 8080,
	}

	conn, err := server.openConn(addr)
	if err != nil {
		t.Fatalf("openConn failed: %v", err)
	}

	if conn == nil {
		t.Fatalf("Expected conn to be non-nil")
	}
}
