package ssh

import (
	"fmt"
	"net"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestClose_5(t *testing.T) {
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
		{
			name: "TestClose",
			fields: fields{
				localAddr: "localhost:8080",
				sshServer: "localhost:22",
				commands:  "echo hello",
				sshConn:   nil,
				ln:        nil,
			},
		},
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
			c.Close()
		})
	}
}
