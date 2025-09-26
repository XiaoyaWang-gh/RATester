package conn

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSetWriteDeadline_1(t *testing.T) {
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
	err := conn.SetWriteDeadline(testTime)

	if err != nil {
		t.Errorf("SetWriteDeadline returned an error: %v", err)
	}
}
