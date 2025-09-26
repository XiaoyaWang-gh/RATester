package vhost

import (
	"fmt"
	"strings"
	"testing"
)

func TestRemoteAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{
		reader: strings.NewReader("Hello, World!"),
	}

	addr := conn.RemoteAddr()
	if addr != nil {
		t.Errorf("Expected RemoteAddr to return nil, but got %v", addr)
	}
}
