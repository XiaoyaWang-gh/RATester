package group

import (
	"fmt"
	"net"
	"testing"

	"gotest.tools/assert"
)

func TestAddr_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ln := &TCPMuxGroupListener{
		groupName: "test",
		group: &TCPMuxGroup{
			group:           "test",
			groupKey:        "test",
			domain:          "test",
			routeByHTTPUser: "test",
			username:        "test",
			acceptCh:        make(chan net.Conn),
		},
		addr: &net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 1234,
		},
	}
	assert.Equal(t, ln.Addr(), ln.addr)
}
