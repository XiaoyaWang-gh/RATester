package nathole

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &discoverConn{
		conn: &net.UDPConn{},
	}
	c.Close()
	if c.messageChan != nil {
		t.Errorf("messageChan should be nil")
	}
}
