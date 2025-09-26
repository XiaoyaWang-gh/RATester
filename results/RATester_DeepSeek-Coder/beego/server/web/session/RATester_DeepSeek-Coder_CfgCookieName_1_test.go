package session

import (
	"fmt"
	"testing"
)

func TestCfgCookieName_1(t *testing.T) {
	type args struct {
		cookieName string
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "Test case 1",
			args: args{
				cookieName: "test_cookie",
			},
			want: func(config *ManagerConfig) {
				config.CookieName = "test_cookie"
			},
		},
		{
			name: "Test case 2",
			args: args{
				cookieName: "another_cookie",
			},
			want: func(config *ManagerConfig) {
				config.CookieName = "another_cookie"
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

			got := CfgCookieName(tt.args.cookieName)
			config := &ManagerConfig{}
			got(config)
			if config.CookieName != tt.args.cookieName {
				t.Errorf("CfgCookieName() = %v, want %v", config.CookieName, tt.args.cookieName)
			}
		})
	}
}
