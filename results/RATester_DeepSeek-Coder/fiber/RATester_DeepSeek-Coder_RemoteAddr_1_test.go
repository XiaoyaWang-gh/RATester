package fiber

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

	conn := &testConn{}
	addr := conn.RemoteAddr()

	if addr.String() != "0.0.0.0:0" {
		t.Errorf("Expected RemoteAddr to be '0.0.0.0:0', got '%s'", addr.String())
	}
}
