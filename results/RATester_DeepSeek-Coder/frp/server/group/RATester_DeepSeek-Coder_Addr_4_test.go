package group

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestAddr_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := &TCPMuxGroupListener{
		groupName: "test",
		addr: &net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8080,
			Zone: "",
		},
		closeCh: make(chan struct{}),
	}

	expected := ln.addr
	actual := ln.Addr()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
