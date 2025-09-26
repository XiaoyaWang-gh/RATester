package net

import (
	"fmt"
	"net"
	"testing"
)

func TestSetRemoteAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &WrapReadWriteCloserConn{}
	addr := &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}
	conn.SetRemoteAddr(addr)

	if conn.remoteAddr != addr {
		t.Errorf("Expected remoteAddr to be %v, but got %v", addr, conn.remoteAddr)
	}
}
