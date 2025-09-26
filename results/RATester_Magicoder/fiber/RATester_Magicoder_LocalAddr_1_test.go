package fiber

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

	conn := &testConn{}
	addr := conn.LocalAddr()
	if addr == nil {
		t.Error("LocalAddr() returned nil")
	}
	if addr.Network() != "tcp" {
		t.Errorf("LocalAddr().Network() = %s, want tcp", addr.Network())
	}
	if addr.String() != "0.0.0.0:0" {
		t.Errorf("LocalAddr().String() = %s, want 0.0.0.0:0", addr.String())
	}
}
