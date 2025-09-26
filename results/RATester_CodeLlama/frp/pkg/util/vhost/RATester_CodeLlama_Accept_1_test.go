package vhost

import (
	"context"
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

	l := &Listener{
		name:            "test",
		location:        "test",
		routeByHTTPUser: "test",
		rewriteHost:     "test",
		username:        "test",
		password:        "test",
		mux:             &Muxer{},
		accept:          make(chan net.Conn),
		ctx:             context.Background(),
	}

	go func() {
		l.accept <- &net.TCPConn{}
	}()

	conn, err := l.Accept()
	if err != nil {
		t.Fatal(err)
	}

	if conn == nil {
		t.Fatal("conn is nil")
	}
}
