package pmux

import (
	"fmt"
	"net"
	"testing"
)

func TestAccept_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pListener := &PortListener{
		connCh: make(chan *PortConn, 1),
	}
	pConn := &PortConn{
		Conn:     &net.TCPConn{},
		rs:       []byte("hello"),
		readMore: true,
	}
	pListener.connCh <- pConn
	conn, err := pListener.Accept()
	if err != nil {
		t.Fatal(err)
	}
	if conn == nil {
		t.Fatal("conn is nil")
	}
	if conn.(*PortConn).Conn != pConn.Conn {
		t.Fatal("conn is not equal")
	}
}
