package file

import (
	"fmt"
	"testing"
)

func TestHasTunnel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		Id: 1,
	}
	tunnel := &Tunnel{
		Port:   8080,
		Client: client,
	}

	exist := client.HasTunnel(tunnel)
	if exist {
		t.Error("Expected false, but got true")
	}

	tunnel.Port = 0
	exist = client.HasTunnel(tunnel)
	if !exist {
		t.Error("Expected true, but got false")
	}
}
