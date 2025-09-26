package net

import (
	"fmt"
	"net"
	"testing"
)

func TestNewFakeUDPConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	laddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}
	raddr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8081}
	l := &UDPListener{
		addr:      laddr,
		acceptCh:  make(chan net.Conn),
		writeCh:   make(chan *UDPPacket),
		readConn:  nil,
		closeFlag: false,
		fakeConns: make(map[string]*FakeUDPConn),
	}

	fc := NewFakeUDPConn(l, laddr, raddr)

	if fc.l != l {
		t.Errorf("Expected l to be %v, but got %v", l, fc.l)
	}

	if fc.localAddr != laddr {
		t.Errorf("Expected localAddr to be %v, but got %v", laddr, fc.localAddr)
	}

	if fc.remoteAddr != raddr {
		t.Errorf("Expected remoteAddr to be %v, but got %v", raddr, fc.remoteAddr)
	}

	if len(fc.packets) != 20 {
		t.Errorf("Expected packets to be of length 20, but got %v", len(fc.packets))
	}

	if fc.closeFlag {
		t.Errorf("Expected closeFlag to be false, but got true")
	}

	if fc.lastActive.IsZero() {
		t.Errorf("Expected lastActive to be non-zero, but got zero")
	}

	fc.Close()

	if !fc.closeFlag {
		t.Errorf("Expected closeFlag to be true, but got false")
	}
}
