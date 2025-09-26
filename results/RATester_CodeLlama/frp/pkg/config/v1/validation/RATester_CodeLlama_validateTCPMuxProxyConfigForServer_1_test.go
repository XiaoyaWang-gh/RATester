package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateTCPMuxProxyConfigForServer_1(t *testing.T) {
	type args struct {
		c *v1.TCPMuxProxyConfig
		s *v1.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				c: &v1.TCPMuxProxyConfig{
					Multiplexer: string(v1.TCPMultiplexerHTTPConnect),
				},
				s: &v1.ServerConfig{
					TCPMuxHTTPConnectPort: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "test case 2",
			args: args{
				c: &v1.TCPMuxProxyConfig{
					Multiplexer: string(v1.TCPMultiplexerHTTPConnect),
				},
				s: &v1.ServerConfig{
					TCPMuxHTTPConnectPort: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := validateTCPMuxProxyConfigForServer(tt.args.c, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("validateTCPMuxProxyConfigForServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
