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
	_, err := conn.Write([]byte("test"))
	if err != io.ErrClosedPipe {
		t.Errorf("Expected ErrClosedPipe, got %v", err)
	}
}
