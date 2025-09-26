package group

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := &TCPMuxGroupListener{
		groupName: "test",
		group:     &TCPMuxGroup{},
		addr:      &net.TCPAddr{},
		closeCh:   make(chan struct{}),
	}

	err := ln.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	select {
	case <-ln.closeCh:
	default:
		t.Errorf("closeCh should be closed")
	}
}
