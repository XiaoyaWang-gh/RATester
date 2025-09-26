package pmux

import (
	"fmt"
	"testing"
)

func TestNewPortMux_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	port := 8080
	managerHost := "127.0.0.1"

	// act
	pMux := NewPortMux(port, managerHost)

	// assert
	if pMux == nil {
		t.Errorf("NewPortMux() failed")
		return
	}
	if pMux.port != port {
		t.Errorf("NewPortMux() failed")
		return
	}
	if pMux.managerHost != managerHost {
		t.Errorf("NewPortMux() failed")
		return
	}
	if pMux.clientConn == nil {
		t.Errorf("NewPortMux() failed")
		return
	}
	if pMux.httpConn == nil {
		t.Errorf("NewPortMux() failed")
		return
	}
	if pMux.httpsConn == nil {
		t.Errorf("NewPortMux() failed")
		return
	}
	if pMux.managerConn == nil {
		t.Errorf("NewPortMux() failed")
		return
	}
}
