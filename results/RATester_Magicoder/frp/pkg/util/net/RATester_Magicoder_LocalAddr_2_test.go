package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestLocalAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &FakeUDPConn{
		localAddr: &net.UDPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 8080,
			Zone: "",
		},
	}

	expected := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
		Zone: "",
	}

	actual := conn.LocalAddr()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
