package ssh

import (
	"fmt"
	"net"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestServeListener_1(t *testing.T) {
	type fields struct {
		localAddr string
		sshServer string
		commands  string
		sshConn   *ssh.Client
		ln        net.Listener
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &TunnelClient{
				localAddr: tt.fields.localAddr,
				sshServer: tt.fields.sshServer,
				commands:  tt.fields.commands,
				sshConn:   tt.fields.sshConn,
				ln:        tt.fields.ln,
			}
			c.serveListener()
		})
	}
}
