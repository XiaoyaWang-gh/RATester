package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgEnableSidInURLQuery_1(t *testing.T) {
	type args struct {
		enable bool
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "Test case 1",
			args: args{
				enable: true,
			},
			want: func(config *ManagerConfig) {
				config.EnableSidInURLQuery = true
			},
		},
		{
			name: "Test case 2",
			args: args{
				enable: false,
			},
			want: func(config *ManagerConfig) {
				config.EnableSidInURLQuery = false
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

			got := CfgEnableSidInURLQuery(tt.args.enable)
			config := &ManagerConfig{}
			got(config)
			wantConfig := &ManagerConfig{}
			tt.want(wantConfig)
			if !reflect.DeepEqual(config, wantConfig) {
				t.Errorf("CfgEnableSidInURLQuery() = %v, want %v", config, wantConfig)
			}
		})
	}
}
