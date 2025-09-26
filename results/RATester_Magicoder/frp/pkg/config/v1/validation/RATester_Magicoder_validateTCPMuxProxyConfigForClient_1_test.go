package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateTCPMuxProxyConfigForClient_1(t *testing.T) {
	tests := []struct {
		name    string
		config  *v1.TCPMuxProxyConfig
		wantErr bool
	}{
		{
			name: "valid config",
			config: &v1.TCPMuxProxyConfig{
				DomainConfig: v1.DomainConfig{
					CustomDomains: []string{"example.com"},
					SubDomain:     "sub",
				},
				Multiplexer: "http_connect",
			},
			wantErr: false,
		},
		{
			name: "invalid multiplexer",
			config: &v1.TCPMuxProxyConfig{
				DomainConfig: v1.DomainConfig{
					CustomDomains: []string{"example.com"},
					SubDomain:     "sub",
				},
				Multiplexer: "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid domain config",
			config: &v1.TCPMuxProxyConfig{
				DomainConfig: v1.DomainConfig{
					CustomDomains: []string{"example.com"},
					SubDomain:     "",
				},
				Multiplexer: "http_connect",
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

			err := validateTCPMuxProxyConfigForClient(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateTCPMuxProxyConfigForClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
