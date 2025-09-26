package common

import (
	"fmt"
	"testing"
)

func TestGetLocalUdpAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn, err := GetLocalUdpAddr()
	if err != nil {
		t.Errorf("GetLocalUdpAddr() error = %v", err)
		return
	}
	defer conn.Close()
	if conn.LocalAddr() == nil {
		t.Errorf("GetLocalUdpAddr() LocalAddr() = nil")
		return
	}
	if conn.RemoteAddr() == nil {
		t.Errorf("GetLocalUdpAddr() RemoteAddr() = nil")
		return
	}
}
