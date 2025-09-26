package conn

import (
	"fmt"
	"net"
	"testing"
)

func TestGetHealthInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: &net.TCPConn{},
		Rb:   []byte("info|status"),
	}

	info, status, err := conn.GetHealthInfo()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if info != "info" {
		t.Errorf("Expected info to be 'info', but got %v", info)
	}

	if status != true {
		t.Errorf("Expected status to be true, but got %v", status)
	}
}
