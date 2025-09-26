package client

import (
	"fmt"
	"net"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestSetInWorkConnCallback_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{}
	cb := func(proxyCfg *v1.ProxyBaseConfig, conn net.Conn, startWorkConn *msg.StartWorkConn) bool {
		return true
	}
	ctl.SetInWorkConnCallback(cb)
}
