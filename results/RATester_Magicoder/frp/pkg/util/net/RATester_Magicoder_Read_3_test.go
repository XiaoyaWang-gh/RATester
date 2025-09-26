package net

import (
	"bytes"
	"fmt"
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
	defer close(conn.packets)

	testData := []byte("test data")
	conn.packets <- testData

	buf := make([]byte, len(testData))
	n, err := conn.Read(buf)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if n != len(testData) {
		t.Errorf("Expected to read %d bytes, but read %d", len(testData), n)
	}
	if !bytes.Equal(buf, testData) {
		t.Errorf("Expected to read %v, but read %v", testData, buf)
	}
}
