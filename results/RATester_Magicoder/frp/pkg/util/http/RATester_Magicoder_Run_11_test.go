package http

import (
	"fmt"
	"net"
	"testing"
)

func TestRun_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		addr: ":8080",
		ln:   &net.TCPListener{},
	}

	err := server.Run()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
