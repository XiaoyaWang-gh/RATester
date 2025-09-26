package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgSessionIdInHTTPHeader_1(t *testing.T) {
	type args struct {
		enable bool
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "TestCfgSessionIdInHTTPHeader_True",
			args: args{enable: true},
			want: func(config *ManagerConfig) {
				config.EnableSidInHTTPHeader = true
			},
		},
		{
			name: "TestCfgSessionIdInHTTPHeader_False",
			args: args{enable: false},
			want: func(config *ManagerConfig) {
				config.EnableSidInHTTPHeader = false
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

			got := CfgSessionIdInHTTPHeader(tt.args.enable)
			config := &ManagerConfig{}
			got(config)
			wantConfig := &ManagerConfig{}
			tt.want(wantConfig)
			if !reflect.DeepEqual(config, wantConfig) {
				t.Errorf("CfgSessionIdInHTTPHeader() = %v, want %v", config, wantConfig)
			}
		})
	}
}
