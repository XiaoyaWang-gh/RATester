package pmux

import (
	"fmt"
	"testing"
)

func TestGetHttpListener_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mux := &PortMux{
		port:     8080,
		httpConn: make(chan *PortConn),
	}

	listener := mux.GetHttpListener()

	if listener == nil {
		t.Error("Expected a listener, but got nil")
	}

	if listener.Addr().String() != ":8080" {
		t.Errorf("Expected listener address to be ':8080', but got '%s'", listener.Addr().String())
	}
}
