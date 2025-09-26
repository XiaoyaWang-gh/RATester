package streamserver

import (
	"fmt"
	"net"
	"testing"
)

func TestRun_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		netType:  TCP,
		bindAddr: "localhost",
		bindPort: 8080,
		handler:  func(c net.Conn) {},
	}

	err := server.Run()
	if err != nil {
		t.Fatalf("Failed to run server: %v", err)
	}

	// Add your test cases here

	server.Close()
}
