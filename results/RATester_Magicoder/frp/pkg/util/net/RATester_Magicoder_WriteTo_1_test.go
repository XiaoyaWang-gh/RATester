package net

import (
	"fmt"
	"net"
	"testing"
)

func TestWriteTo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	testConn := &ConnectedUDPConn{conn.(*net.UDPConn)}

	testData := []byte("Hello, World!")
	n, err := testConn.WriteTo(testData, conn.RemoteAddr())
	if err != nil {
		t.Fatal(err)
	}
	if n != len(testData) {
		t.Fatalf("Expected to write %d bytes, but wrote %d", len(testData), n)
	}
}
