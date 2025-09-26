package ssh

import (
	"fmt"
	"net"
	"testing"
)

func TestRun_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &Gateway{
		bindPort:           8080,
		ln:                 nil,
		peerServerListener: nil,
		sshConfig:          nil,
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", g.bindPort))
	if err != nil {
		t.Fatalf("Failed to listen on port %d: %v", g.bindPort, err)
	}
	g.ln = ln

	go g.Run()

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", g.bindPort))
	if err != nil {
		t.Fatalf("Failed to dial port %d: %v", g.bindPort, err)
	}
	conn.Close()

	g.ln.Close()
}
