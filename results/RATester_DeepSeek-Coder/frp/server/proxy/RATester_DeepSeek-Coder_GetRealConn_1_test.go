package proxy

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestGetRealConn_1(t *testing.T) {
	type fields struct {
		BaseProxy  *BaseProxy
		cfg        *v1.HTTPProxyConfig
		closeFuncs []func()
	}
	type args struct {
		remoteAddr string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantWorkConn net.Conn
		wantErr      bool
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

			pxy := &HTTPProxy{
				BaseProxy:  tt.fields.BaseProxy,
				cfg:        tt.fields.cfg,
				closeFuncs: tt.fields.closeFuncs,
			}
			gotWorkConn, err := pxy.GetRealConn(tt.args.remoteAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPProxy.GetRealConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotWorkConn, tt.wantWorkConn) {
				t.Errorf("HTTPProxy.GetRealConn() = %v, want %v", gotWorkConn, tt.wantWorkConn)
			}
		})
	}
}
