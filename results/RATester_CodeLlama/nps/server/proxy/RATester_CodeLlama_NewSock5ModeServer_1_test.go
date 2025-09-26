package proxy

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestNewSock5ModeServer_1(t *testing.T) {
	type args struct {
		bridge NetBridge
		task   *file.Tunnel
	}
	tests := []struct {
		name string
		args args
		want *Sock5ModeServer
	}{
		{
			name: "test",
			args: args{
				bridge: nil,
				task:   nil,
			},
			want: &Sock5ModeServer{
				BaseServer: BaseServer{
					id:           -1,
					bridge:       nil,
					task:         nil,
					errorContent: nil,
					Mutex:        sync.Mutex{},
				},
				listener: nil,
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

			if got := NewSock5ModeServer(tt.args.bridge, tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSock5ModeServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
