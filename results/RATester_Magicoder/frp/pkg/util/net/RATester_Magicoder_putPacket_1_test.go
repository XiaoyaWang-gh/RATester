package net

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPutPacket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &FakeUDPConn{
		packets: make(chan []byte, 1),
	}

	content := []byte("test content")
	conn.putPacket(content)

	select {
	case packet := <-conn.packets:
		if !bytes.Equal(packet, content) {
			t.Errorf("Expected packet to be %v, but got %v", content, packet)
		}
	default:
		t.Error("Expected packet to be put in the channel, but it was not")
	}
}
