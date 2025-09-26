package bridge

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/file"
)

func TestNewTunnel_1(t *testing.T) {
	type args struct {
		tunnelPort     int
		tunnelType     string
		ipVerify       bool
		runList        sync.Map
		disconnectTime int
	}
	tests := []struct {
		name string
		args args
		want *Bridge
	}{
		{
			name: "test1",
			args: args{
				tunnelPort:     10000,
				tunnelType:     "tcp",
				ipVerify:       true,
				runList:        sync.Map{},
				disconnectTime: 10,
			},
			want: &Bridge{
				TunnelPort:     10000,
				tunnelType:     "tcp",
				OpenTask:       make(chan *file.Tunnel),
				CloseTask:      make(chan *file.Tunnel),
				CloseClient:    make(chan int),
				SecretChan:     make(chan *conn.Secret),
				ipVerify:       true,
				runList:        sync.Map{},
				disconnectTime: 10,
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

			if got := NewTunnel(tt.args.tunnelPort, tt.args.tunnelType, tt.args.ipVerify, tt.args.runList, tt.args.disconnectTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTunnel() = %v, want %v", got, tt.want)
			}
		})
	}
}
