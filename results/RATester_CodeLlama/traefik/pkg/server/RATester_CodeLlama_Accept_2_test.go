package server

import (
	"fmt"
	"net"
	"testing"
)

func TestAccept_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := tcpKeepAliveListener{
		TCPListener: &net.TCPListener{},
	}

	conn, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}

	if conn == nil {
		t.Fatal("conn is nil")
	}
}
