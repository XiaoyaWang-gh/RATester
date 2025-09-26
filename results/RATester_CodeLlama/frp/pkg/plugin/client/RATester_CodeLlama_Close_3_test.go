package client

import (
	"fmt"
	"net"
	"testing"
)

func TestClose_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	uds := &UnixDomainSocketPlugin{
		UnixAddr: &net.UnixAddr{},
	}
	err := uds.Close()
	if err != nil {
		t.Error(err)
	}
}
