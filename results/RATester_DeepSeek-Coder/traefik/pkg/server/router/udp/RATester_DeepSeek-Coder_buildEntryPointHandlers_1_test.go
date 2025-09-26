package udp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
	"github.com/traefik/traefik/v3/pkg/udp"
)

func TestBuildEntryPointHandlers_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		configs map[string]*runtime.UDPRouterInfo
	}
	tests := []struct {
		name string
		args args
		want []udp.Handler
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
			if got := m.buildEntryPointHandlers(tt.args.ctx, tt.args.configs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.buildEntryPointHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
