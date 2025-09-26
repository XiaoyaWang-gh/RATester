package net

import (
	"fmt"
	"testing"
	"time"
)

func TestSetWriteDeadline_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &FakeUDPConn{}

	testTime := time.Now()
	err := conn.SetWriteDeadline(testTime)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
