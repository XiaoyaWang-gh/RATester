package net

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetDeadline_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &FakeUDPConn{
		localAddr: &net.UDPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8080,
			Zone: "",
		},
	}

	testTime := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
	err := conn.SetDeadline(testTime)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if conn.lastActive != testTime {
		t.Errorf("Expected lastActive to be %v, got %v", testTime, conn.lastActive)
	}
}
