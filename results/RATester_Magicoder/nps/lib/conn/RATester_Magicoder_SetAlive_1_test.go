package conn

import (
	"fmt"
	"net"
	"testing"

	"ehang.io/nps/lib/pmux"
	"github.com/xtaci/kcp-go"
)

func TestSetAlive_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: &kcp.UDPSession{},
	}
	conn.SetAlive("")

	conn = &Conn{
		Conn: &net.TCPConn{},
	}
	conn.SetAlive("")

	conn = &Conn{
		Conn: &pmux.PortConn{},
	}
	conn.SetAlive("")
}
