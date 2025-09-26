package pmux

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetWriteDeadline_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		pConn := &PortConn{Conn: conn}
		deadline := time.Now().Add(1 * time.Second)
		err = pConn.SetWriteDeadline(deadline)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1)
	_, err = conn.Read(buf)
	if err != nil {
		t.Error(err)
	}
}
