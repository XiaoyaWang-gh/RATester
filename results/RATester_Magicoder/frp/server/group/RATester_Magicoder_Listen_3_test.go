package group

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestListen_3(t *testing.T) {
	type args struct {
		ctx         context.Context
		multiplexer string
		group       string
		groupKey    string
		routeConfig vhost.RouteConfig
	}
	tests := []struct {
		name    string
		args    args
		wantL   net.Listener
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

			tmgc := &TCPMuxGroupCtl{}
			gotL, err := tmgc.Listen(tt.args.ctx, tt.args.multiplexer, tt.args.group, tt.args.groupKey, tt.args.routeConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("TCPMuxGroupCtl.Listen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotL, tt.wantL) {
				t.Errorf("TCPMuxGroupCtl.Listen() = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}
