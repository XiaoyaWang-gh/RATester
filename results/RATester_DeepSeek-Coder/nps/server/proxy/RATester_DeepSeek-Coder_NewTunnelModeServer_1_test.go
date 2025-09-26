package proxy

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestNewTunnelModeServer_1(t *testing.T) {
	type args struct {
		process process
		bridge  NetBridge
		task    *file.Tunnel
	}
	tests := []struct {
		name string
		args args
		want *TunnelModeServer
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

			if got := NewTunnelModeServer(tt.args.process, tt.args.bridge, tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTunnelModeServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
