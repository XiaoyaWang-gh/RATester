package proxy

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestHttpConnectRun_1(t *testing.T) {
	type fields struct {
		BaseProxy *BaseProxy
		cfg       *v1.TCPMuxProxyConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
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

			pxy := &TCPMuxProxy{
				BaseProxy: tt.fields.BaseProxy,
				cfg:       tt.fields.cfg,
			}
			got, err := pxy.httpConnectRun()
			if (err != nil) != tt.wantErr {
				t.Errorf("httpConnectRun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("httpConnectRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
