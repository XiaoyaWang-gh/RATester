package ssh

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/fatedier/frp/client/proxy"
)

func TestWaitProxyStatusReady_1(t *testing.T) {
	type args struct {
		name    string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		s       *TunnelServer
		args    args
		want    *proxy.WorkingStatus
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

			got, err := tt.s.waitProxyStatusReady(tt.args.name, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("TunnelServer.waitProxyStatusReady() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TunnelServer.waitProxyStatusReady() = %v, want %v", got, tt.want)
			}
		})
	}
}
