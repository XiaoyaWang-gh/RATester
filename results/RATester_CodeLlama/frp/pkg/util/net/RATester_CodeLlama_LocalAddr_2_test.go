package net

import (
	"fmt"
	"net"
	"testing"

	"gotest.tools/assert"
)

func TestLocalAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &FakeUDPConn{
		localAddr: &net.UDPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 1234,
		},
	}
	assert.Equal(t, c.LocalAddr(), c.localAddr)
}
