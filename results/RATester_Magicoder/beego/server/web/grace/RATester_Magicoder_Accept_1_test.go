package grace

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

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Error(err)
		}
		defer conn.Close()

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Error(err)
		}
		if n != 1024 {
			t.Errorf("Expected to read 1024 bytes, but read %d", n)
		}
	}()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write(make([]byte, 1024))
	if err != nil {
		t.Error(err)
	}
}
