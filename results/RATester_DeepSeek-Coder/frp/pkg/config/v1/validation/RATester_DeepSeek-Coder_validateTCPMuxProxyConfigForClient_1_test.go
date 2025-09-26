package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateTCPMuxProxyConfigForClient_1(t *testing.T) {
	type args struct {
		c *v1.TCPMuxProxyConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				c: &v1.TCPMuxProxyConfig{
					DomainConfig: v1.DomainConfig{
						CustomDomains: []string{"example.com"},
						SubDomain:     "sub",
					},
					Multiplexer: "TCPMultiplexerHTTPConnect",
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				c: &v1.TCPMuxProxyConfig{
					DomainConfig: v1.DomainConfig{
						CustomDomains: []string{"example.com"},
						SubDomain:     "sub",
					},
					Multiplexer: "TCPMultiplexerUnsupported",
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

			if err := validateTCPMuxProxyConfigForClient(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateTCPMuxProxyConfigForClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
