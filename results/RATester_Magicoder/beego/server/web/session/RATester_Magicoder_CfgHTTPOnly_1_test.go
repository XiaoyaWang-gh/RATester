package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgHTTPOnly_1(t *testing.T) {
	type args struct {
		HTTPOnly bool
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "TestCfgHTTPOnly_True",
			args: args{HTTPOnly: true},
			want: func(config *ManagerConfig) {
				config.DisableHTTPOnly = false
			},
		},
		{
			name: "TestCfgHTTPOnly_False",
			args: args{HTTPOnly: false},
			want: func(config *ManagerConfig) {
				config.DisableHTTPOnly = true
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

			got := CfgHTTPOnly(tt.args.HTTPOnly)
			config := &ManagerConfig{}
			got(config)
			wantConfig := &ManagerConfig{}
			tt.want(wantConfig)
			if !reflect.DeepEqual(config, wantConfig) {
				t.Errorf("CfgHTTPOnly() = %v, want %v", config, wantConfig)
			}
		})
	}
}
