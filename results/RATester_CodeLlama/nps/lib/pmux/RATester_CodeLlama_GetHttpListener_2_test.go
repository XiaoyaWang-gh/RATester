package pmux

import (
	"fmt"
	"net"
	"testing"
)

func TestGetHttpListener_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pMux := &PortMux{}
	pMux.httpConn = make(chan *PortConn, 1)
	pMux.clientConn = make(chan *PortConn, 1)
	pMux.httpsConn = make(chan *PortConn, 1)
	pMux.managerConn = make(chan *PortConn, 1)
	pMux.port = 8080
	pMux.isClose = false
	pMux.managerHost = "127.0.0.1"
	pMux.Listener = &net.TCPListener{}
	pMux.Listener.Addr()
	pMux.GetHttpListener()
}
