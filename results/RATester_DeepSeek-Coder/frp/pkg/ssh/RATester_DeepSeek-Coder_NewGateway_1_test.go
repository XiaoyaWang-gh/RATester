package ssh

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestNewGateway_1(t *testing.T) {
	type args struct {
		cfg                v1.SSHTunnelGateway
		bindAddr           string
		peerServerListener *netpkg.InternalListener
	}
	tests := []struct {
		name    string
		args    args
		want    *Gateway
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

			got, err := NewGateway(tt.args.cfg, tt.args.bindAddr, tt.args.peerServerListener)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGateway() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGateway() = %v, want %v", got, tt.want)
			}
		})
	}
}
