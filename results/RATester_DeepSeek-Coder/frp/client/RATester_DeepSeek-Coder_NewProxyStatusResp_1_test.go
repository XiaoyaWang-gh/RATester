package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/client/proxy"
)

func TestNewProxyStatusResp_1(t *testing.T) {
	type args struct {
		status     *proxy.WorkingStatus
		serverAddr string
	}
	tests := []struct {
		name string
		args args
		want ProxyStatusResp
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

			if got := NewProxyStatusResp(tt.args.status, tt.args.serverAddr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxyStatusResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
