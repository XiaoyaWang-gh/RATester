package conn

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetReadDeadline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: &net.TCPConn{},
		Rb:   []byte{},
	}

	testTime := time.Now()
	err := conn.SetReadDeadline(testTime)

	if err != nil {
		t.Errorf("SetReadDeadline returned an error: %v", err)
	}
}
