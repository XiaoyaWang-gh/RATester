package pmux

import (
	"fmt"
	"testing"
)

func TestClose_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	pMux := &PortMux{
		port: 8080,
	}
	// act
	err := pMux.Close()
	// assert
	if err != nil {
		t.Errorf("Close() error = %v", err)
		return
	}
	if pMux.isClose != true {
		t.Errorf("Close() isClose = %v, want %v", pMux.isClose, true)
		return
	}
	if pMux.clientConn != nil {
		t.Errorf("Close() clientConn = %v, want nil", pMux.clientConn)
		return
	}
	if pMux.httpConn != nil {
		t.Errorf("Close() httpConn = %v, want nil", pMux.httpConn)
		return
	}
	if pMux.httpsConn != nil {
		t.Errorf("Close() httpsConn = %v, want nil", pMux.httpsConn)
		return
	}
	if pMux.managerConn != nil {
		t.Errorf("Close() managerConn = %v, want nil", pMux.managerConn)
		return
	}
}
