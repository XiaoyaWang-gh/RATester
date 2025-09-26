package pmux

import (
	"fmt"
	"testing"
)

func TestStart_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pMux := &PortMux{
		port: 80,
	}
	err := pMux.Start()
	if err != nil {
		t.Error(err)
	}
	if pMux.port != 80 {
		t.Error("port not equal")
	}
	if pMux.Listener == nil {
		t.Error("Listener is nil")
	}
	if pMux.clientConn == nil {
		t.Error("clientConn is nil")
	}
	if pMux.httpConn == nil {
		t.Error("httpConn is nil")
	}
	if pMux.httpsConn == nil {
		t.Error("httpsConn is nil")
	}
	if pMux.managerConn == nil {
		t.Error("managerConn is nil")
	}
	pMux.Close()
}
