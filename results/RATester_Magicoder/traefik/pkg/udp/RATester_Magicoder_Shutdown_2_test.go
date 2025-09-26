package udp

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

func TestShutdown_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{
		pConn:     &net.UDPConn{},
		mu:        sync.RWMutex{},
		conns:     make(map[string]*Conn),
		accepting: true,
		acceptCh:  make(chan *Conn),
		timeout:   10 * time.Second,
	}

	graceTimeout := 5 * time.Second

	err := l.Shutdown(graceTimeout)
	if err != nil {
		t.Errorf("Shutdown returned an error: %v", err)
	}

	if l.accepting {
		t.Errorf("Shutdown did not set accepting to false")
	}

	if len(l.conns) != 0 {
		t.Errorf("Shutdown did not close all connections")
	}
}
