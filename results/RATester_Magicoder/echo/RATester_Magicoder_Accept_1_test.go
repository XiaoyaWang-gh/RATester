package echo

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
		conn, err := net.Dial("tcp", ln.Addr().String())
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write([]byte("Hello, world!"))
		if err != nil {
			t.Error(err)
			return
		}
	}()

	conn, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 14)
	_, err = conn.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if string(buf) != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!', got %s", string(buf))
	}
}
