package net

import (
	"fmt"
	"testing"
	"time"
)

func TestSetReadDeadline_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fakeUDPConn := &FakeUDPConn{}
	fakeUDPConn.SetReadDeadline(time.Now())
	if fakeUDPConn.IsClosed() {
		t.Error("fakeUDPConn should not be closed")
	}
}
