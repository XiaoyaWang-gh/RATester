package udp

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestReadLoop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		receiveCh: make(chan []byte),
		readCh:    make(chan []byte),
		sizeCh:    make(chan int),
		msgs:      make([][]byte, 0),
		timeout:   10 * time.Second,
		doneCh:    make(chan struct{}),
	}

	go conn.readLoop()

	msg := []byte("test message")
	conn.receiveCh <- msg

	buf := make([]byte, len(msg))
	conn.readCh <- buf

	n := <-conn.sizeCh
	if n != len(msg) {
		t.Errorf("Expected to read %d bytes, but read %d", len(msg), n)
	}

	if !bytes.Equal(buf[:n], msg) {
		t.Errorf("Expected to read %v, but read %v", msg, buf[:n])
	}

	conn.close()
}
