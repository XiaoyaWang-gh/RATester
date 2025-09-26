package ssh

import (
	"fmt"
	"net"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestStart_3(t *testing.T) {
	type fields struct {
		localAddr string
		sshServer string
		commands  string
		sshConn   *ssh.Client
		ln        net.Listener
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test case 1",
			fields: fields{
				localAddr: "127.0.0.1:80",
				sshServer: "127.0.0.1:22",
				commands:  "command",
				sshConn:   nil,
				ln:        nil,
			},
			wantErr: false,
		},
		// Add more test cases here
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
			if err := c.Start(); (err != nil) != tt.wantErr {
				t.Errorf("TunnelClient.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
