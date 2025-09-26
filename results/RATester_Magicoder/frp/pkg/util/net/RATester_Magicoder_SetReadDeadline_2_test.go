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

	conn := &FakeUDPConn{}
	now := time.Now()
	err := conn.SetReadDeadline(now)
	if err != nil {
		t.Errorf("SetReadDeadline returned an error: %v", err)
	}
}
