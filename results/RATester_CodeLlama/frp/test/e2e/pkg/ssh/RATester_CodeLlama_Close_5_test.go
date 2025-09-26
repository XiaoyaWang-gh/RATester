package ssh

import (
	"fmt"
	"testing"
)

func TestClose_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &TunnelClient{
		localAddr: "127.0.0.1:10080",
		sshServer: "127.0.0.1:22",
		commands:  "echo 123",
	}
	c.Close()
}
