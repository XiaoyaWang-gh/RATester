package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestSendHealthInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: &net.TCPConn{},
		Rb:   []byte{},
	}

	info := "testInfo"
	status := "testStatus"

	_, err := conn.SendHealthInfo(info, status)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
