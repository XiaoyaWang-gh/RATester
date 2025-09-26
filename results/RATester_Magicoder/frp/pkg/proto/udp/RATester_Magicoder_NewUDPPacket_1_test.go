package udp

import (
	"encoding/base64"
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewUDPPacket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	buf := []byte("test data")
	laddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	raddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8081,
	}

	packet := NewUDPPacket(buf, laddr, raddr)

	if packet.Content != base64.StdEncoding.EncodeToString(buf) {
		t.Errorf("Expected Content to be %s, but got %s", base64.StdEncoding.EncodeToString(buf), packet.Content)
	}

	if !reflect.DeepEqual(packet.LocalAddr, laddr) {
		t.Errorf("Expected LocalAddr to be %v, but got %v", laddr, packet.LocalAddr)
	}

	if !reflect.DeepEqual(packet.RemoteAddr, raddr) {
		t.Errorf("Expected RemoteAddr to be %v, but got %v", raddr, packet.RemoteAddr)
	}
}
