package tcp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetHTTPRouters_2(t *testing.T) {
	type args struct {
		ctx         context.Context
		entryPoints []string
		tls         bool
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]*runtime.RouterInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Manager{}
			if got := m.getHTTPRouters(tt.args.ctx, tt.args.entryPoints, tt.args.tls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.getHTTPRouters() = %v, want %v", got, tt.want)
			}
		})
	}
}
