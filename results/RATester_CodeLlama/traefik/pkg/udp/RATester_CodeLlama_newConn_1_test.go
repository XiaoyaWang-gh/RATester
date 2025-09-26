package udp

import (
	"fmt"
	"net"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	l := &Listener{}
	rAddr := &net.UDPAddr{}

	// when
	conn := l.newConn(rAddr)

	// then
	assert.NotNil(t, conn)
	assert.Equal(t, l, conn.listener)
	assert.Equal(t, rAddr, conn.rAddr)
	assert.NotNil(t, conn.receiveCh)
	assert.NotNil(t, conn.readCh)
	assert.NotNil(t, conn.sizeCh)
	assert.NotNil(t, conn.doneCh)
	assert.Equal(t, l.timeout, conn.timeout)
}
