package net

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestRead_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &FakeUDPConn{
		packets: make(chan []byte, 1),
	}

	testData := []byte("test data")
	conn.packets <- testData

	b := make([]byte, len(testData))
	n, err := conn.Read(b)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if n != len(testData) {
		t.Errorf("Expected to read %d bytes, got %d", len(testData), n)
	}
	if !bytes.Equal(b, testData) {
		t.Errorf("Expected to read %v, got %v", testData, b)
	}

	close(conn.packets)

	n, err = conn.Read(b)
	if err != io.EOF {
		t.Errorf("Expected EOF error, got %v", err)
	}
	if n != 0 {
		t.Errorf("Expected to read 0 bytes, got %d", n)
	}
}
