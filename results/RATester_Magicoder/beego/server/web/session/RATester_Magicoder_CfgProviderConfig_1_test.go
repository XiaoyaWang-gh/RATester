package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgProviderConfig_1(t *testing.T) {
	type args struct {
		providerConfig string
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "Test case 1",
			args: args{
				providerConfig: "test_config",
			},
			want: func(config *ManagerConfig) {
				config.ProviderConfig = "test_config"
			},
		},
		{
			name: "Test case 2",
			args: args{
				providerConfig: "another_config",
			},
			want: func(config *ManagerConfig) {
				config.ProviderConfig = "another_config"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CfgProviderConfig(tt.args.providerConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CfgProviderConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
