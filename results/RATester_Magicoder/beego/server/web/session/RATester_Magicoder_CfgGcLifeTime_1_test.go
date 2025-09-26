package session

import (
	"fmt"
	"reflect"
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
			name: "Test case 1",
			args: args{lifeTime: 10},
			want: func(config *ManagerConfig) {
				config.Gclifetime = 10
			},
		},
		{
			name: "Test case 2",
			args: args{lifeTime: 20},
			want: func(config *ManagerConfig) {
				config.Gclifetime = 20
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

			if got := CfgGcLifeTime(tt.args.lifeTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CfgGcLifeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
