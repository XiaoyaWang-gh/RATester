package nathole

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn, err := net.Dial("udp", "localhost:1234")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	dc := &discoverConn{
		conn:        conn.(*net.UDPConn),
		localAddr:   conn.LocalAddr(),
		messageChan: make(chan *Message),
	}

	err = dc.Close()
	if err != nil {
		t.Fatal(err)
	}

	if dc.messageChan != nil {
		t.Error("messageChan should be nil after Close")
	}

	if dc.conn != nil {
		t.Error("conn should be nil after Close")
	}
}
