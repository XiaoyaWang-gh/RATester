package net

import (
	"fmt"
	"net"
	"testing"
)

func TestAddr_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := &WebsocketListener{
		ln:       &net.TCPListener{},
		acceptCh: make(chan net.Conn),
	}
	addr := ln.Addr()
	if addr == nil {
		t.Fatal("addr is nil")
	}
}
