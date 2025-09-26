package grace

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestAccept_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := tcpKeepAliveListener{
		TCPListener: &net.TCPListener{},
	}
	c, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	if c == nil {
		t.Fatal("c is nil")
	}
	if c.LocalAddr() == nil {
		t.Fatal("c.LocalAddr() is nil")
	}
	if c.RemoteAddr() == nil {
		t.Fatal("c.RemoteAddr() is nil")
	}
	if c.SetDeadline(time.Time{}) != nil {
		t.Fatal("c.SetDeadline() failed")
	}
	if c.SetReadDeadline(time.Time{}) != nil {
		t.Fatal("c.SetReadDeadline() failed")
	}
	if c.SetWriteDeadline(time.Time{}) != nil {
		t.Fatal("c.SetWriteDeadline() failed")
	}
}
