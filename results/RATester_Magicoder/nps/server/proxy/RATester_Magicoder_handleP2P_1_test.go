package proxy

import (
	"fmt"
	"net"
	"testing"

	"ehang.io/nps/lib/common"
)

func TesthandleP2P_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testServer := &P2PServer{
		p2p: make(map[string]*p2p),
	}

	testAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}

	testStr := "test_role" + common.CONN_DATA_SEQ + "test_password"

	testServer.handleP2P(testAddr, testStr)

	// Add assertions here to check the state of testServer after the call to handleP2P
}
