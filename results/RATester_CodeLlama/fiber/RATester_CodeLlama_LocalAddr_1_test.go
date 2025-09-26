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
	if addr.Network() != "tcp" {
		t.Errorf("LocalAddr().Network() = %q; want %q", addr.Network(), "tcp")
	}
	if addr.String() != "0.0.0.0:0" {
		t.Errorf("LocalAddr().String() = %q; want %q", addr.String(), "0.0.0.0:0")
	}
}
