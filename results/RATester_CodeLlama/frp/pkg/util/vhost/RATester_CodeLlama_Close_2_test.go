package vhost

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestClose_2(t *testing.T) {
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
	err := l.Close()
	if err != nil {
		t.Error(err)
	}
}
