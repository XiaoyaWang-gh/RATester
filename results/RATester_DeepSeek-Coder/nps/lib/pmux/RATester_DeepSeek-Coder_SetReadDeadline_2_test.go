package pmux

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetReadDeadline_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTime := time.Now()
	portConn := &PortConn{
		Conn: &net.TCPConn{},
	}

	err := portConn.SetReadDeadline(testTime)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
