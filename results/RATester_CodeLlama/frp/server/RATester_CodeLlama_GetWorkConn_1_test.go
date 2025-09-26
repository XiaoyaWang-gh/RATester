package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/zeebo/assert"
)

func TestGetWorkConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{
		workConnCh: make(chan net.Conn, 1),
	}

	workConn, err := ctl.GetWorkConn()
	assert.NoError(t, err)
	assert.NotNil(t, workConn)

	workConn, err = ctl.GetWorkConn()
	assert.NoError(t, err)
	assert.NotNil(t, workConn)

	close(ctl.workConnCh)
	workConn, err = ctl.GetWorkConn()
	assert.Error(t, err)
	assert.Nil(t, workConn)
}
