package common

import (
	"fmt"
	"testing"
)

func TestNewUDPDatagram_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	header := &UDPHeader{
		Rsv:  1,
		Frag: 1,
		Addr: &Addr{},
	}
	data := []byte("test data")
	datagram := NewUDPDatagram(header, data)

	if datagram.Header.Rsv != header.Rsv {
		t.Errorf("Expected Rsv to be %d, but got %d", header.Rsv, datagram.Header.Rsv)
	}

	if datagram.Header.Frag != header.Frag {
		t.Errorf("Expected Frag to be %d, but got %d", header.Frag, datagram.Header.Frag)
	}

	if datagram.Header.Addr != header.Addr {
		t.Errorf("Expected Addr to be %p, but got %p", header.Addr, datagram.Header.Addr)
	}

	if string(datagram.Data) != string(data) {
		t.Errorf("Expected Data to be %s, but got %s", string(data), string(datagram.Data))
	}
}
