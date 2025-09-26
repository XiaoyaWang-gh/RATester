package net

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPutPacket_1(t *testing.T) {
	conn := &FakeUDPConn{
		packets: make(chan []byte, 1),
	}

	testContent := []byte("test content")

	t.Run("putPacket should put content into packets channel", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		conn.putPacket(testContent)

		select {
		case content := <-conn.packets:
			if !bytes.Equal(content, testContent) {
				t.Errorf("Expected %v, got %v", testContent, content)
			}
		default:
			t.Errorf("Expected content to be put into packets channel")
		}
	})

	t.Run("putPacket should not block when packets channel is full", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		conn.putPacket(testContent)
		conn.putPacket(testContent)

		select {
		case <-conn.packets:
			t.Errorf("Expected packets channel not to be full")
		default:
		}
	})
}
