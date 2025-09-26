package udp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestGetUDPRouters_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		entryPoints []string
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]*runtime.UDPRouterInfo
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
			if got := m.getUDPRouters(tt.args.ctx, tt.args.entryPoints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.getUDPRouters() = %v, want %v", got, tt.want)
			}
		})
	}
}
