package proxy

import (
	"fmt"
	"net"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestRun_14(t *testing.T) {
	type fields struct {
		BaseProxy *BaseProxy
		cfg       *v1.UDPProxyConfig
		localAddr *net.UDPAddr
		readCh    chan *msg.UDPPacket
		sendCh    chan msg.Message
		workConn  net.Conn
		closed    bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
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

			pxy := &UDPProxy{
				BaseProxy: tt.fields.BaseProxy,
				cfg:       tt.fields.cfg,
				localAddr: tt.fields.localAddr,
				readCh:    tt.fields.readCh,
				sendCh:    tt.fields.sendCh,
				workConn:  tt.fields.workConn,
				closed:    tt.fields.closed,
			}
			if err := pxy.Run(); (err != nil) != tt.wantErr {
				t.Errorf("UDPProxy.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
