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

	c := &FakeUDPConn{}
	if c.IsClosed() {
		t.Error("IsClosed() should return false when the connection is not closed")
	}
	c.Close()
	if !c.IsClosed() {
		t.Error("IsClosed() should return true when the connection is closed")
	}
}
