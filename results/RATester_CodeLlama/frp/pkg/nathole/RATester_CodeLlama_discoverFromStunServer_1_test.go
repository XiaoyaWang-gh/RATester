package nathole

import (
	"fmt"
	"testing"
)

func TestDiscoverFromStunServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	conn := &discoverConn{}
	addr := "127.0.0.1:19302"

	// when
	externalAddrs, err := conn.discoverFromStunServer(addr)

	// then
	if err != nil {
		t.Errorf("discoverFromStunServer() error = %v", err)
		return
	}
	if len(externalAddrs) != 1 {
		t.Errorf("discoverFromStunServer() externalAddrs = %v, want %v", externalAddrs, 1)
		return
	}
	if externalAddrs[0] != "127.0.0.1:19302" {
		t.Errorf("discoverFromStunServer() externalAddrs = %v, want %v", externalAddrs, "127.0.0.1:19302")
		return
	}
}
