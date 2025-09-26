package nathole

import (
	"fmt"
	"net"
	"testing"
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

	d := &discoverConn{
		conn:        conn,
		localAddr:   conn.LocalAddr(),
		messageChan: make(chan *Message, 1),
	}

	go d.readLoop()

	msg := []byte("Hello, world")
	_, err = conn.WriteToUDP(msg, &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	if err != nil {
		t.Fatal(err)
	}

	received := <-d.messageChan
	if string(received.Body) != string(msg) {
		t.Errorf("Expected message: %s, got: %s", msg, received.Body)
	}
	if received.Addr != conn.LocalAddr().String() {
		t.Errorf("Expected address: %s, got: %s", conn.LocalAddr().String(), received.Addr)
	}
}
