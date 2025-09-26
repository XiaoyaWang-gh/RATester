package vhost

import (
	"fmt"
	"testing"
)

func TestLocalAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := readOnlyConn{reader: nil}
	if conn.LocalAddr() != nil {
		t.Errorf("LocalAddr() = %v, want nil", conn.LocalAddr())
	}
}
