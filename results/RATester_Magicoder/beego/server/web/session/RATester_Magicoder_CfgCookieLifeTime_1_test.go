package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgCookieLifeTime_1(t *testing.T) {
	type args struct {
		lifeTime int
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "TestCfgCookieLifeTime",
			args: args{lifeTime: 10},
			want: func(config *ManagerConfig) {
				config.CookieLifeTime = 10
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

			if got := CfgCookieLifeTime(tt.args.lifeTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CfgCookieLifeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
