package net

import (
	"fmt"
	"testing"
)

func TestIsClosed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &FakeUDPConn{
		closeFlag: false,
	}

	if conn.IsClosed() {
		t.Error("Expected IsClosed() to return false, but got true")
	}

	conn.closeFlag = true
	if !conn.IsClosed() {
		t.Error("Expected IsClosed() to return true, but got false")
	}
}
