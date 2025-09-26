package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOpts_1(t *testing.T) {
	type args struct {
		opts []ManagerConfigOpt
	}
	tests := []struct {
		name string
		c    *ManagerConfig
		args args
		want *ManagerConfig
	}{
		{
			name: "TestOpts",
			c:    &ManagerConfig{},
			args: args{
				opts: []ManagerConfigOpt{
					func(c *ManagerConfig) {
						c.EnableSetCookie = true
					},
					func(c *ManagerConfig) {
						c.DisableHTTPOnly = true
					},
				},
			},
			want: &ManagerConfig{
				EnableSetCookie:         true,
				DisableHTTPOnly:         true,
				Secure:                  false,
				EnableSidInHTTPHeader:   false,
				EnableSidInURLQuery:     false,
				CookieName:              "",
				Gclifetime:              0,
				Maxlifetime:             0,
				CookieLifeTime:          0,
				ProviderConfig:          "",
				Domain:                  "",
				SessionIDLength:         0,
				SessionNameInHTTPHeader: "",
				SessionIDPrefix:         "",
				CookieSameSite:          0,
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

			tt.c.Opts(tt.args.opts...)
			if !reflect.DeepEqual(tt.c, tt.want) {
				t.Errorf("ManagerConfig.Opts() = %v, want %v", tt.c, tt.want)
			}
		})
	}
}
