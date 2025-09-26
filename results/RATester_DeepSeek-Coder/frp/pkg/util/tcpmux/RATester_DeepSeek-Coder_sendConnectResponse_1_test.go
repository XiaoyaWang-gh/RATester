package tcpmux

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestSendConnectResponse_1(t *testing.T) {
	type fields struct {
		Muxer       *vhost.Muxer
		passthrough bool
	}
	type args struct {
		c net.Conn
		m map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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

			muxer := &HTTPConnectTCPMuxer{
				Muxer:       tt.fields.Muxer,
				passthrough: tt.fields.passthrough,
			}
			if err := muxer.sendConnectResponse(tt.args.c, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("HTTPConnectTCPMuxer.sendConnectResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
