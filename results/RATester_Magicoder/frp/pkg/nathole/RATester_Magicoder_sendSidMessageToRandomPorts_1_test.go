package nathole

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestSendSidMessageToRandomPorts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	addrs := []string{"127.0.0.1"}
	count := 10
	sendFunc := func(c *net.UDPConn, addr string) error {
		_, err := c.WriteToUDP([]byte("test"), &net.UDPAddr{IP: net.ParseIP(addr)})
		return err
	}

	sendSidMessageToRandomPorts(ctx, conn, addrs, count, sendFunc)

	// Add assertions here to verify the behavior of the function under test.
}
