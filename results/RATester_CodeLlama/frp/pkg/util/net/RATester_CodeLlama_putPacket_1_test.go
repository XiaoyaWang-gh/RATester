package net

import (
	"fmt"
	"testing"
)

func TestPutPacket_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &FakeUDPConn{}
	c.packets = make(chan []byte, 1)
	c.putPacket([]byte("test"))
	if len(c.packets) != 1 {
		t.Fatal("putPacket failed")
	}
}
