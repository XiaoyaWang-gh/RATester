package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateTCPProxyConfigForServer_1(t *testing.T) {
	type args struct {
		c *v1.TCPProxyConfig
		s *v1.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
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

			if err := validateTCPProxyConfigForServer(tt.args.c, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("validateTCPProxyConfigForServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
