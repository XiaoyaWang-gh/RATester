package proxy

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"ehang.io/nps/bridge"
	"ehang.io/nps/lib/file"
)

func TestNewBaseServer_1(t *testing.T) {
	type args struct {
		bridge *bridge.Bridge
		task   *file.Tunnel
	}
	tests := []struct {
		name string
		args args
		want *BaseServer
	}{
		{
			name: "test",
			args: args{
				bridge: &bridge.Bridge{},
				task:   &file.Tunnel{},
			},
			want: &BaseServer{
				bridge:       &bridge.Bridge{},
				task:         &file.Tunnel{},
				errorContent: nil,
				Mutex:        sync.Mutex{},
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

			if got := NewBaseServer(tt.args.bridge, tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
