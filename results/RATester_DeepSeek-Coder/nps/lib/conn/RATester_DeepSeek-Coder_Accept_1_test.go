package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestAccept_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		Accept(listener, func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1024)
			n, err := c.Read(buf)
			if err != nil {
				t.Error(err)
			}
			if n != 1024 {
				t.Errorf("expected to read 1024 bytes, got %d", n)
			}
		})
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write(make([]byte, 1024))
	if err != nil {
		t.Fatal(err)
	}

	listener.Close()
	<-done
}
