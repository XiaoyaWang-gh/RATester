package streamserver

import (
	"fmt"
	"net"
	"testing"
)

func TestHandle_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		netType:     TCP,
		bindAddr:    "127.0.0.1",
		bindPort:    8080,
		respContent: []byte("test response"),
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", server.bindAddr, server.bindPort))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("test request"))
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if string(buf[:n]) != string(server.respContent) {
		t.Fatalf("expected response %s, got %s", string(server.respContent), string(buf[:n]))
	}
}
