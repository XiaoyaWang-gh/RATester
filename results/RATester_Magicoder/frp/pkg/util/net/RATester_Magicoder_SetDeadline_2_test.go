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

	conn := &FakeUDPConn{}
	now := time.Now()
	err := conn.SetDeadline(now)
	if err != nil {
		t.Errorf("SetDeadline returned an error: %v", err)
	}
}
