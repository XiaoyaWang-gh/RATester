package client

import (
	"fmt"
	"testing"
)

func TestnewUdpConn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &TRPClient{
		svrAddr:        "127.0.0.1:8080",
		bridgeConnType: "tcp",
	}

	localAddr := "127.0.0.1:8081"
	rAddr := "127.0.0.1:8082"
	md5Password := "password"

	client.newUdpConn(localAddr, rAddr, md5Password)

	// Add assertions here to check the expected behavior
}
