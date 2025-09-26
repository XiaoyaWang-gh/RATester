package net

import (
	"fmt"
	"testing"
	"time"
)

func TestSetDeadline_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fakeUDPConn := &FakeUDPConn{}
	fakeUDPConn.SetDeadline(time.Now())
	if fakeUDPConn.IsClosed() {
		t.Errorf("fakeUDPConn should not be closed")
	}
}
