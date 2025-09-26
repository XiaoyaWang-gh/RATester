package echo

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

	ln := tcpKeepAliveListener{}
	c, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	if err := c.(*net.TCPConn).SetKeepAlive(true); err != nil {
		t.Fatal(err)
	}

	if err := c.(*net.TCPConn).SetKeepAlivePeriod(3 * time.Minute); err != nil {
		t.Fatal(err)
	}
}
