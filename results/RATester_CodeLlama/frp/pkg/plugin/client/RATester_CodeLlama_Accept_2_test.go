package client

import (
	"fmt"
	"net"
	"testing"

	"github.com/zeebo/assert"
)

func TestAccept_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{
		conns: make(chan net.Conn, 1),
	}
	conn := &net.TCPConn{}
	l.conns <- conn
	c, err := l.Accept()
	assert.NoError(t, err)
	assert.Equal(t, conn, c)
}
