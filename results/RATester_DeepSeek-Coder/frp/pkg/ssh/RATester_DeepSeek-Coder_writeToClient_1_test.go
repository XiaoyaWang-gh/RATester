package ssh

import (
	"fmt"
	"net"
	"sync"
	"testing"

	netpkg "github.com/fatedier/frp/pkg/util/net"
	"github.com/fatedier/frp/pkg/virtual"
	"golang.org/x/crypto/ssh"
)

func TestWriteToClient_1(t *testing.T) {
	type fields struct {
		underlyingConn     net.Conn
		sshConn            *ssh.ServerConn
		sc                 *ssh.ServerConfig
		firstChannel       ssh.Channel
		vc                 *virtual.Client
		peerServerListener *netpkg.InternalListener
		doneCh             chan struct{}
		closeDoneChOnce    sync.Once
	}
	type args struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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

			s := &TunnelServer{
				underlyingConn:     tt.fields.underlyingConn,
				sshConn:            tt.fields.sshConn,
				sc:                 tt.fields.sc,
				firstChannel:       tt.fields.firstChannel,
				vc:                 tt.fields.vc,
				peerServerListener: tt.fields.peerServerListener,
				doneCh:             tt.fields.doneCh,
				closeDoneChOnce:    tt.fields.closeDoneChOnce,
			}
			s.writeToClient(tt.args.data)
			// TODO: check the result
		})
	}
}
