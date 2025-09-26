package conn

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetDeadline_1(t *testing.T) {
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
	err := conn.SetDeadline(testTime)

	if err != nil {
		t.Errorf("SetDeadline returned an error: %v", err)
	}
}
