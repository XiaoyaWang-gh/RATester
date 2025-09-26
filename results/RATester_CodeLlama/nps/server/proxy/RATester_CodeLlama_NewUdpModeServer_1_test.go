package proxy

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"ehang.io/nps/bridge"
	"ehang.io/nps/lib/file"
)

func TestNewUdpModeServer_1(t *testing.T) {
	type args struct {
		bridge *bridge.Bridge
		task   *file.Tunnel
	}
	tests := []struct {
		name string
		args args
		want *UdpModeServer
	}{
		{
			name: "test1",
			args: args{
				bridge: &bridge.Bridge{},
				task:   &file.Tunnel{},
			},
			want: &UdpModeServer{
				BaseServer: BaseServer{
					id:           0,
					bridge:       &bridge.Bridge{},
					task:         &file.Tunnel{},
					errorContent: []byte{},
					Mutex:        sync.Mutex{},
				},
				addrMap: sync.Map{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewUdpModeServer(tt.args.bridge, tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUdpModeServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
