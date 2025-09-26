package session

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewManagerConfig_1(t *testing.T) {
	tests := []struct {
		name string
		want *ManagerConfig
	}{
		{
			name: "Test case 1",
			want: &ManagerConfig{
				EnableSetCookie:       true,
				DisableHTTPOnly:       false,
				Secure:                true,
				EnableSidInHTTPHeader: true,
				EnableSidInURLQuery:   true,
				CookieName:            "session",
				Gclifetime:            3600,
				Maxlifetime:           3600,
				CookieLifeTime:        3600,
				ProviderConfig:        "redis",
				Domain:                "localhost",
				SessionIDLength:       32,
				SessionIDPrefix:       "sess_",
				CookieSameSite:        http.SameSiteDefaultMode,
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

			if got := NewManagerConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManagerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
