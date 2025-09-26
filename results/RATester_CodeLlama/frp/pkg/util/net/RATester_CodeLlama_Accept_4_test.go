package net

import (
	"fmt"
	"net"
	"testing"
)

func TestAccept_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := &WebsocketListener{
		ln:       nil,
		acceptCh: make(chan net.Conn, 1),
		server:   nil,
	}
	ln.acceptCh <- &net.TCPConn{}
	conn, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	if conn == nil {
		t.Fatal("conn is nil")
	}
}
