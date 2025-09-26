package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgMaxLifeTime_1(t *testing.T) {
	type args struct {
		lifeTime int64
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "TestCfgMaxLifeTime",
			args: args{lifeTime: 10},
			want: func(config *ManagerConfig) {
				config.Maxlifetime = 10
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

			if got := CfgMaxLifeTime(tt.args.lifeTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CfgMaxLifeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
