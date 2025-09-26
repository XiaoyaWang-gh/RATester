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

	if addr.Network() != "" {
		t.Errorf("Expected Network() to return an empty string, got %s", addr.Network())
	}

	if addr.String() != "0.0.0.0:0" {
		t.Errorf("Expected String() to return '0.0.0.0:0', got %s", addr.String())
	}
}
