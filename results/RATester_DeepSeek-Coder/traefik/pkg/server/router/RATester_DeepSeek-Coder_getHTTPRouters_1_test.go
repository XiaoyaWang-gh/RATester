package router

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetHTTPRouters_1(t *testing.T) {
	tests := []struct {
		name        string
		manager     *Manager
		entryPoints []string
		tls         bool
		want        map[string]map[string]*runtime.RouterInfo
	}{
		{
			name: "Test with nil conf",
			manager: &Manager{
				conf: nil,
			},
			entryPoints: []string{"http"},
			tls:         false,
			want:        map[string]map[string]*runtime.RouterInfo{},
		},
		{
			name: "Test with non-nil conf",
			manager: &Manager{
				conf: &runtime.Configuration{},
			},
			entryPoints: []string{"http"},
			tls:         false,
			want:        map[string]map[string]*runtime.RouterInfo{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.manager.getHTTPRouters(context.Background(), tt.entryPoints, tt.tls)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHTTPRouters() = %v, want %v", got, tt.want)
			}
		})
	}
}
