package client

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	uds := &UnixDomainSocketPlugin{}
	if uds.Name() != v1.PluginUnixDomainSocket {
		t.Fatal("Name() failed")
	}
}
