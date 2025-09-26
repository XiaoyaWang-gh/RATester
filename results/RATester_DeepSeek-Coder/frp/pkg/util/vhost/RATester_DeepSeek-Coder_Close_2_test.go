package vhost

import (
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
		location:        "localhost:8080",
		routeByHTTPUser: "route",
		mux:             &Muxer{},
		accept:          make(chan net.Conn),
	}

	err := l.Close()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	select {
	case <-l.accept:
		t.Errorf("Expected closed channel, got open channel")
	default:
	}
}
