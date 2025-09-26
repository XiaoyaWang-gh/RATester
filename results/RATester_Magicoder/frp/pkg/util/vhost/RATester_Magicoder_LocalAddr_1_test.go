package vhost

import (
	"fmt"
	"strings"
	"testing"
)

func TestLocalAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{
		reader: strings.NewReader("test data"),
	}

	addr := conn.LocalAddr()
	if addr != nil {
		t.Errorf("Expected LocalAddr to return nil, but got %v", addr)
	}
}
