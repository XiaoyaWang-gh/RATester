package group

import (
	"fmt"
	"net"
	"testing"
)

func TestWorker_4(t *testing.T) {
	type fields struct {
		group    string
		groupKey string
		domain   string
		acceptCh chan net.Conn
		tcpMuxLn net.Listener
		lns      []*TCPMuxGroupListener
		ctl      *TCPMuxGroupCtl
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

			tmg := &TCPMuxGroup{
				group:    tt.fields.group,
				groupKey: tt.fields.groupKey,
				domain:   tt.fields.domain,
				acceptCh: tt.fields.acceptCh,
				tcpMuxLn: tt.fields.tcpMuxLn,
				lns:      tt.fields.lns,
				ctl:      tt.fields.ctl,
			}
			tmg.worker()
		})
	}
}
