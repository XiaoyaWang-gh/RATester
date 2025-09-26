package udp

import (
	"encoding/base64"
	"fmt"
	"net"
	"testing"

	"gotest.tools/assert"
)

func TestNewUDPPacket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := []byte("hello world")
	laddr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 1234}
	raddr := &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 5678}
	packet := NewUDPPacket(buf, laddr, raddr)
	assert.Equal(t, base64.StdEncoding.EncodeToString(buf), packet.Content)
	assert.Equal(t, laddr, packet.LocalAddr)
	assert.Equal(t, raddr, packet.RemoteAddr)
}
