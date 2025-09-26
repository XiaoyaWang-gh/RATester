package vhost

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestRun_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	muxer := &Muxer{
		listener: &net.TCPListener{},
	}

	go muxer.run()

	// Test case 1: Accept returns an error
	muxer.listener.(*net.TCPListener).SetDeadline(time.Now().Add(1 * time.Second))
	_, err := muxer.listener.Accept()
	if err == nil {
		t.Error("Expected an error but got nil")
	}

	// Test case 2: Accept returns a connection
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	muxer.handle(conn)

	// Test case 3: Handle function is called
	muxer.handle(conn)

	// Test case 4: Handle function is not called
	muxer.handle(nil)
}
