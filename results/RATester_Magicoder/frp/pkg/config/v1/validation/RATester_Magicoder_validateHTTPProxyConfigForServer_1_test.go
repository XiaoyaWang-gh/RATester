package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateHTTPProxyConfigForServer_1(t *testing.T) {
	type args struct {
		c *v1.HTTPProxyConfig
		s *v1.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				c: &v1.HTTPProxyConfig{
					DomainConfig: v1.DomainConfig{
						CustomDomains: []string{"example.com"},
					},
				},
				s: &v1.ServerConfig{
					VhostHTTPPort: 80,
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				c: &v1.HTTPProxyConfig{
					DomainConfig: v1.DomainConfig{
						CustomDomains: []string{"example.com"},
					},
				},
				s: &v1.ServerConfig{
					VhostHTTPPort: 0,
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

			if err := validateHTTPProxyConfigForServer(tt.args.c, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("validateHTTPProxyConfigForServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
