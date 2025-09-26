package ssh

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestParseClientAndProxyConfigurer_1(t *testing.T) {
	type args struct {
		extraPayload string
	}
	tests := []struct {
		name    string
		s       *TunnelServer
		args    args
		want    *v1.ClientCommonConfig
		want1   v1.ProxyConfigurer
		want2   string
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

			got, got1, got2, err := tt.s.parseClientAndProxyConfigurer(nil, tt.args.extraPayload)
			if (err != nil) != tt.wantErr {
				t.Errorf("TunnelServer.parseClientAndProxyConfigurer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TunnelServer.parseClientAndProxyConfigurer() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("TunnelServer.parseClientAndProxyConfigurer() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("TunnelServer.parseClientAndProxyConfigurer() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
