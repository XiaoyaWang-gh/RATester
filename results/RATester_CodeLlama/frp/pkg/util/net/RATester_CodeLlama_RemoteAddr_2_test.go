package net

import (
	"fmt"
	"net"
	"testing"

	"gotest.tools/assert"
)

func TestRemoteAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &FakeUDPConn{
		remoteAddr: &net.UDPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 12345,
		},
	}
	assert.Equal(t, c.RemoteAddr(), c.remoteAddr)
}
