package ssh

import (
	"fmt"
	"testing"
)

func TestNewTunnelClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	localAddr := "localhost:8080"
	sshServer := "localhost:22"
	commands := "ls"

	client := NewTunnelClient(localAddr, sshServer, commands)

	if client.localAddr != localAddr {
		t.Errorf("Expected localAddr to be %s, but got %s", localAddr, client.localAddr)
	}

	if client.sshServer != sshServer {
		t.Errorf("Expected sshServer to be %s, but got %s", sshServer, client.sshServer)
	}

	if client.commands != commands {
		t.Errorf("Expected commands to be %s, but got %s", commands, client.commands)
	}

	if client.sshConn != nil {
		t.Errorf("Expected sshConn to be nil, but got %v", client.sshConn)
	}

	if client.ln != nil {
		t.Errorf("Expected ln to be nil, but got %v", client.ln)
	}
}
