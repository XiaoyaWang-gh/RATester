package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateSUDPProxyConfigForServer_1(t *testing.T) {
	type args struct {
		c *v1.SUDPProxyConfig
		s *v1.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				c: &v1.SUDPProxyConfig{
					ProxyBaseConfig: v1.ProxyBaseConfig{
						Name: "test1",
					},
				},
				s: &v1.ServerConfig{
					BindAddr: "127.0.0.1",
				},
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				c: &v1.SUDPProxyConfig{
					ProxyBaseConfig: v1.ProxyBaseConfig{
						Name: "test2",
					},
				},
				s: &v1.ServerConfig{
					BindAddr: "127.0.0.1",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := validateSUDPProxyConfigForServer(tt.args.c, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("validateSUDPProxyConfigForServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
