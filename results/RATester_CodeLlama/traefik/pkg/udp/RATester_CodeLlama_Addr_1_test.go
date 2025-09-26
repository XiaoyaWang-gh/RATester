package udp

import (
	"fmt"
	"net"
	"testing"
)

func TestAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{}
	l.pConn = &net.UDPConn{}
	l.pConn.LocalAddr()
	l.Addr()
}
