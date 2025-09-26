package virtual

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/fatedier/frp/client"
	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestPeerListener_1(t *testing.T) {
	type fields struct {
		l   *netpkg.InternalListener
		svr *client.Service
	}
	tests := []struct {
		name   string
		fields fields
		want   net.Listener
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

			c := &Client{
				l:   tt.fields.l,
				svr: tt.fields.svr,
			}
			if got := c.PeerListener(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.PeerListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
