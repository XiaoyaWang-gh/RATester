package nathole

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestReadLoop_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	dc := &discoverConn{
		conn:        conn,
		localAddr:   conn.LocalAddr(),
		messageChan: make(chan *Message, 1),
	}

	go dc.readLoop()

	msg := []byte("Hello, world!")
	_, err = conn.WriteToUDP(msg, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	if err != nil {
		t.Fatal(err)
	}

	select {
	case m := <-dc.messageChan:
		if !bytes.Equal(m.Body, msg) {
			t.Errorf("Expected message body %v, got %v", msg, m.Body)
		}
		if m.Addr != "127.0.0.1:0" {
			t.Errorf("Expected message addr 127.0.0.1:0, got %v", m.Addr)
		}
	case <-time.After(time.Second):
		t.Error("Timed out waiting for message")
	}
}
