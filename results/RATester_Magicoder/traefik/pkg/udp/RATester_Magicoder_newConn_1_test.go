package udp

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNewConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{
		conns:   make(map[string]*Conn),
		timeout: 10 * time.Second,
	}
	rAddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}
	conn := l.newConn(rAddr)

	if conn.listener != l {
		t.Errorf("Expected listener to be %v, but got %v", l, conn.listener)
	}

	if conn.rAddr != rAddr {
		t.Errorf("Expected rAddr to be %v, but got %v", rAddr, conn.rAddr)
	}

	if conn.timeout != 10*time.Second {
		t.Errorf("Expected timeout to be 10s, but got %v", conn.timeout)
	}

	if conn.receiveCh == nil {
		t.Error("Expected receiveCh to be initialized, but it is nil")
	}

	if conn.readCh == nil {
		t.Error("Expected readCh to be initialized, but it is nil")
	}

	if conn.sizeCh == nil {
		t.Error("Expected sizeCh to be initialized, but it is nil")
	}

	if conn.doneCh == nil {
		t.Error("Expected doneCh to be initialized, but it is nil")
	}
}
