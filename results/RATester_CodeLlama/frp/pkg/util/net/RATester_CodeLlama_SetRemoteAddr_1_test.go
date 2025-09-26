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
	addr := &net.TCPAddr{}
	conn.SetRemoteAddr(addr)
	if conn.remoteAddr != addr {
		t.Errorf("conn.remoteAddr = %v, want %v", conn.remoteAddr, addr)
	}
}
