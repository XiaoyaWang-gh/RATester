package bridge

import (
	"fmt"
	"net"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestGetHealthFromClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bridge := &Bridge{
		TunnelPort: 8080,
	}

	conn := &conn.Conn{
		Conn: &net.TCPConn{},
	}

	bridge.GetHealthFromClient(1, conn)

	// Add assertions here to verify the behavior of your method.
}
