package session

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCfgSameSite_1(t *testing.T) {
	type args struct {
		sameSite http.SameSite
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "TestCfgSameSite_Lax",
			args: args{
				sameSite: http.SameSiteLaxMode,
			},
			want: func(config *ManagerConfig) {
				config.CookieSameSite = http.SameSiteLaxMode
			},
		},
		{
			name: "TestCfgSameSite_Strict",
			args: args{
				sameSite: http.SameSiteStrictMode,
			},
			want: func(config *ManagerConfig) {
				config.CookieSameSite = http.SameSiteStrictMode
			},
		},
		{
			name: "TestCfgSameSite_None",
			args: args{
				sameSite: http.SameSiteNoneMode,
			},
			want: func(config *ManagerConfig) {
				config.CookieSameSite = http.SameSiteNoneMode
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

			got := CfgSameSite(tt.args.sameSite)
			config := &ManagerConfig{}
			got(config)
			if config.CookieSameSite != tt.args.sameSite {
				t.Errorf("CfgSameSite() = %v, want %v", config.CookieSameSite, tt.args.sameSite)
			}
		})
	}
}
