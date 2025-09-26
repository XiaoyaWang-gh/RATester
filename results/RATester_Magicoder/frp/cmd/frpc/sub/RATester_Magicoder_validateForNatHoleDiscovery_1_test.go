package sub

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateForNatHoleDiscovery_1(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *v1.ClientCommonConfig
		wantErr bool
	}{
		{
			name: "test case 1",
			cfg: &v1.ClientCommonConfig{
				NatHoleSTUNServer: "",
			},
			wantErr: true,
		},
		{
			name: "test case 2",
			cfg: &v1.ClientCommonConfig{
				NatHoleSTUNServer: "stun.example.com",
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

			if err := validateForNatHoleDiscovery(tt.cfg); (err != nil) != tt.wantErr {
				t.Errorf("validateForNatHoleDiscovery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
