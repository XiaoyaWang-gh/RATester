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

	_, err = conn.Write([]byte("Hello, world!"))
	if err != nil {
		t.Errorf("conn.Write() error = %v", err)
	}
}
