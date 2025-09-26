package nathole

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestSendSidMessageToRangePorts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	addrs := []string{"127.0.0.1"}
	ports := []msg.PortsRange{{From: 8080, To: 8081}}

	sendFunc := func(c *net.UDPConn, addr string) error {
		_, err := c.Write([]byte("test"))
		return err
	}

	sendSidMessageToRangePorts(ctx, conn.(*net.UDPConn), addrs, ports, sendFunc)
}
