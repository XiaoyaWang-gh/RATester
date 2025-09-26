package pmux

import (
	"fmt"
	"net"
	"testing"
)

func TestGetHttpsListener_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pMux := &PortMux{}
	pMux.httpsConn = make(chan *PortConn, 1)
	pMux.Listener = &net.TCPListener{}
	pMux.port = 8080
	pMux.managerHost = "127.0.0.1"
	pMux.clientConn = make(chan *PortConn, 1)
	pMux.httpConn = make(chan *PortConn, 1)
	pMux.managerConn = make(chan *PortConn, 1)
	pMux.isClose = false
	pMux.Start()
	pMux.GetHttpsListener()
	if pMux.GetHttpsListener() == nil {
		t.Errorf("GetHttpsListener() = %v, want %v", pMux.GetHttpsListener(), nil)
	}
}
