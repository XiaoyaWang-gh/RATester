package visitor

import (
	"fmt"
	"net"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestDispatcher_1(t *testing.T) {
	type fields struct {
		BaseVisitor  *BaseVisitor
		checkCloseCh chan struct{}
		udpConn      *net.UDPConn
		readCh       chan *msg.UDPPacket
		sendCh       chan *msg.UDPPacket
		cfg          *v1.SUDPVisitorConfig
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

			sv := &SUDPVisitor{
				BaseVisitor:  tt.fields.BaseVisitor,
				checkCloseCh: tt.fields.checkCloseCh,
				udpConn:      tt.fields.udpConn,
				readCh:       tt.fields.readCh,
				sendCh:       tt.fields.sendCh,
				cfg:          tt.fields.cfg,
			}
			sv.dispatcher()
		})
	}
}
