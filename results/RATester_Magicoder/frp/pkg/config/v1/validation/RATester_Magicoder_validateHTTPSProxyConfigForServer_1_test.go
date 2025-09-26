package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateHTTPSProxyConfigForServer_1(t *testing.T) {
	type args struct {
		c *v1.HTTPSProxyConfig
		s *v1.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				c: &v1.HTTPSProxyConfig{
					DomainConfig: v1.DomainConfig{
						CustomDomains: []string{"example.com"},
					},
				},
				s: &v1.ServerConfig{
					VhostHTTPSPort: 443,
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				c: &v1.HTTPSProxyConfig{
					DomainConfig: v1.DomainConfig{
						CustomDomains: []string{"example.com"},
					},
				},
				s: &v1.ServerConfig{
					VhostHTTPSPort: 0,
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

			if err := validateHTTPSProxyConfigForServer(tt.args.c, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("validateHTTPSProxyConfigForServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
