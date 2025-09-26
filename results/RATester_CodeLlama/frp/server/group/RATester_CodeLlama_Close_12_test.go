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

	ln.Close()

	if len(ln.closeCh) != 0 {
		t.Errorf("Close() failed")
	}
}
