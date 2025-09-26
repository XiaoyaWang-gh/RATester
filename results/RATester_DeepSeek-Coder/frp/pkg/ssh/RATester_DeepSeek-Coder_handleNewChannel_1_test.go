package ssh

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestHandleNewChannel_1(t *testing.T) {
	type args struct {
		channel        ssh.NewChannel
		extraPayloadCh chan string
	}
	tests := []struct {
		name string
		s    *TunnelServer
		args args
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

			tt.s.handleNewChannel(tt.args.channel, tt.args.extraPayloadCh)
		})
	}
}
