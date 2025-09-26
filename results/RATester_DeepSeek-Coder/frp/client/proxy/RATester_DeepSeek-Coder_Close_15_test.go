package proxy

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestClose_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &UDPProxy{
		closed: false,
		readCh: make(chan *msg.UDPPacket),
	}

	pxy.Close()

	if !pxy.closed {
		t.Errorf("Expected closed to be true after calling Close, but it was false")
	}

	_, ok := <-pxy.readCh
	if ok {
		t.Errorf("Expected readCh to be closed after calling Close, but it was still open")
	}
}
