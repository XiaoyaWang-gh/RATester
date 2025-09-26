package sub

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateForNatHoleDiscovery_1(t *testing.T) {
	type args struct {
		cfg *v1.ClientCommonConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				cfg: &v1.ClientCommonConfig{
					NatHoleSTUNServer: "stun.stunprotocol.org",
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				cfg: &v1.ClientCommonConfig{
					NatHoleSTUNServer: "",
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

			if err := validateForNatHoleDiscovery(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("validateForNatHoleDiscovery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
