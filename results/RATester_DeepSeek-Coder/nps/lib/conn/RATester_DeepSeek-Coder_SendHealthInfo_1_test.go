package conn

import (
	"fmt"
	"testing"
)

func TestSendHealthInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: nil,
		Rb:   []byte{},
	}

	info := "test info"
	status := "test status"

	n, err := conn.SendHealthInfo(info, status)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if n == 0 {
		t.Errorf("Expected to write more than 0 bytes, got %v", n)
	}
}
