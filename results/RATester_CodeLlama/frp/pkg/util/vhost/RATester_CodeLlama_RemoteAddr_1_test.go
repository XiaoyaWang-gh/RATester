package vhost

import (
	"fmt"
	"testing"
)

func TestRemoteAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{reader: nil}
	if conn.RemoteAddr() != nil {
		t.Errorf("conn.RemoteAddr() = %v, want nil", conn.RemoteAddr())
	}
}
