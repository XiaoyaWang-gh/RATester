package udp

import (
	"fmt"
	"net"
	"testing"
)

func TestGetConn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{
		conns:     make(map[string]*Conn),
		accepting: true,
		acceptCh:  make(chan *Conn),
	}

	raddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}

	conn, err := l.getConn(raddr)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if conn == nil {
		t.Fatalf("Expected conn, got nil")
	}

	if _, ok := l.conns[raddr.String()]; !ok {
		t.Fatalf("Expected conn to be in conns map")
	}

	if _, ok := <-l.acceptCh; !ok {
		t.Fatalf("Expected conn to be sent on acceptCh")
	}
}
