package client

import (
	"fmt"
	"net"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	uds := &UnixDomainSocketPlugin{
		UnixAddr: &net.UnixAddr{},
	}

	expected := v1.PluginUnixDomainSocket
	actual := uds.Name()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
