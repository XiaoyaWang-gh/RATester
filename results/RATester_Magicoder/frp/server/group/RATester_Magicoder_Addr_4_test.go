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
		groupName: "testGroup",
		group:     &TCPMuxGroup{},
		addr:      &net.TCPAddr{},
		closeCh:   make(chan struct{}),
	}

	expected := ln.addr
	actual := ln.Addr()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
