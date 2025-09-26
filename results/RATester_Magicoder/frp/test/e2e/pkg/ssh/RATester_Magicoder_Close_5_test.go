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

	client := &TunnelClient{
		localAddr: "localhost:8080",
		sshServer: "localhost:22",
	}

	client.Close()

	if client.sshConn != nil || client.ln != nil {
		t.Error("sshConn and ln should be nil after Close()")
	}
}
