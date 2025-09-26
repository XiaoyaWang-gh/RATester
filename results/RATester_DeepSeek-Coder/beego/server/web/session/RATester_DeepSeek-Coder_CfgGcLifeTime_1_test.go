package session

import (
	"fmt"
	"testing"
)

func TestCfgGcLifeTime_1(t *testing.T) {
	type args struct {
		lifeTime int64
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "TestCfgGcLifeTime",
			args: args{lifeTime: 10},
			want: func(config *ManagerConfig) {
				config.Gclifetime = 10
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

			got := CfgGcLifeTime(tt.args.lifeTime)
			config := &ManagerConfig{}
			got(config)
			if config.Gclifetime != tt.args.lifeTime {
				t.Errorf("CfgGcLifeTime() = %v, want %v", config.Gclifetime, tt.args.lifeTime)
			}
		})
	}
}
