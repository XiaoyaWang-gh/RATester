package pmux

import (
	"fmt"
	"testing"
)

func TestGetClientListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mux := &PortMux{
		port:        8080,
		isClose:     false,
		managerHost: "localhost",
		clientConn:  make(chan *PortConn),
		httpConn:    make(chan *PortConn),
		httpsConn:   make(chan *PortConn),
		managerConn: make(chan *PortConn),
	}

	listener := mux.GetClientListener()
	if listener == nil {
		t.Error("Expected a listener, but got nil")
	}

	mux.Close()
	listener = mux.GetClientListener()
	if listener != nil {
		t.Error("Expected nil listener after closing, but got a listener")
	}
}
