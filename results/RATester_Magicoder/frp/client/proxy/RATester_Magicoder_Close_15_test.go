package proxy

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/golib/msg/json"
)

func TestClose_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &UDPProxy{
		closed:   false,
		workConn: &net.UDPConn{},
		readCh:   make(chan *msg.UDPPacket),
		sendCh:   make(chan json.Message),
	}

	pxy.Close()

	if pxy.closed != true {
		t.Error("pxy.closed should be true")
	}

	if pxy.workConn != nil {
		t.Error("pxy.workConn should be nil")
	}

	if pxy.readCh != nil {
		t.Error("pxy.readCh should be nil")
	}

	if pxy.sendCh != nil {
		t.Error("pxy.sendCh should be nil")
	}
}
