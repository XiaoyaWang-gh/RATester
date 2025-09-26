package net

import (
	"fmt"
	"net"
	"testing"

	"github.com/zeebo/assert"
)

func TestPutConn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &InternalListener{
		acceptCh: make(chan net.Conn, 1),
	}
	conn := &net.TCPConn{}
	err := l.PutConn(conn)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(l.acceptCh))
	assert.Equal(t, conn, <-l.acceptCh)
}
