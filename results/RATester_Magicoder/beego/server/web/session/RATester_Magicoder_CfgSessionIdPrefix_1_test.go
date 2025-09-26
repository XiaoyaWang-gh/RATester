package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgSessionIdPrefix_1(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "Test case 1",
			args: args{
				prefix: "test_prefix",
			},
			want: func(config *ManagerConfig) {
				config.SessionIDPrefix = "test_prefix"
			},
		},
		{
			name: "Test case 2",
			args: args{
				prefix: "another_prefix",
			},
			want: func(config *ManagerConfig) {
				config.SessionIDPrefix = "another_prefix"
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

			got := CfgSessionIdPrefix(tt.args.prefix)
			config := &ManagerConfig{}
			got(config)
			wantConfig := &ManagerConfig{}
			tt.want(wantConfig)
			if !reflect.DeepEqual(config, wantConfig) {
				t.Errorf("CfgSessionIdPrefix() = %v, want %v", config, wantConfig)
			}
		})
	}
}
