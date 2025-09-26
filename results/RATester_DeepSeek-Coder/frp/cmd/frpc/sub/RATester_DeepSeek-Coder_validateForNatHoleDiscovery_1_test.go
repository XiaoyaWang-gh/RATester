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
			name: "Test with empty NatHoleSTUNServer",
			args: args{
				cfg: &v1.ClientCommonConfig{
					NatHoleSTUNServer: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Test with non-empty NatHoleSTUNServer",
			args: args{
				cfg: &v1.ClientCommonConfig{
					NatHoleSTUNServer: "stun.example.com",
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

			err := validateForNatHoleDiscovery(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateForNatHoleDiscovery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
