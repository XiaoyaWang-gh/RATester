package vhost

import (
	"fmt"
	"io"
	"testing"
)

func TestWrite_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{}
	n, err := conn.Write([]byte("test"))
	if n != 0 {
		t.Errorf("Expected 0 bytes written, got %d", n)
	}
	if err != io.ErrClosedPipe {
		t.Errorf("Expected error %v, got %v", io.ErrClosedPipe, err)
	}
}
