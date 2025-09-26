package tcp

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestAddTCPHandlers_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		configs map[string]*runtime.TCPRouterInfo
		router  *Router
	}
	tests := []struct {
		name string
		args args
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
			m.addTCPHandlers(tt.args.ctx, tt.args.configs, tt.args.router)
		})
	}
}
