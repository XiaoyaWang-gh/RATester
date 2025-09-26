package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateXTCPProxyConfigForClient_1(t *testing.T) {
	type args struct {
		c *v1.XTCPProxyConfig
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

			if err := validateXTCPProxyConfigForClient(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateXTCPProxyConfigForClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
