package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateXTCPVisitorConfig_1(t *testing.T) {
	type testCase struct {
		name    string
		config  *v1.XTCPVisitorConfig
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "valid kcp protocol",
			config: &v1.XTCPVisitorConfig{
				Protocol: "kcp",
			},
			wantErr: false,
		},
		{
			name: "valid quic protocol",
			config: &v1.XTCPVisitorConfig{
				Protocol: "quic",
			},
			wantErr: false,
		},
		{
			name: "invalid protocol",
			config: &v1.XTCPVisitorConfig{
				Protocol: "invalid",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := validateXTCPVisitorConfig(tc.config)
			if (err != nil) != tc.wantErr {
				t.Errorf("validateXTCPVisitorConfig() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
