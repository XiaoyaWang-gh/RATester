package udp

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRead_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		readCh: make(chan []byte),
		sizeCh: make(chan int),
	}

	data := []byte("test data")
	go func() {
		conn.readCh <- data
		conn.sizeCh <- len(data)
	}()

	buf := make([]byte, len(data))
	n, err := conn.Read(buf)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if n != len(data) {
		t.Errorf("Expected to read %d bytes, but read %d", len(data), n)
	}

	if !bytes.Equal(buf, data) {
		t.Errorf("Expected to read %v, but read %v", data, buf)
	}
}
