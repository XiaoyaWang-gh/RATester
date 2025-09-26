package visitor

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewQUICTunnelSession_1(t *testing.T) {
	type args struct {
		clientCfg *v1.ClientCommonConfig
	}
	tests := []struct {
		name string
		args args
		want TunnelSession
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

			if got := NewQUICTunnelSession(tt.args.clientCfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQUICTunnelSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
