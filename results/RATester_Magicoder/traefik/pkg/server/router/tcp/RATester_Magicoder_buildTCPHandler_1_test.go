package tcp

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestBuildTCPHandler_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		router *runtime.TCPRouterInfo
	}
	tests := []struct {
		name    string
		args    args
		want    tcp.Handler
		wantErr bool
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
			got, err := m.buildTCPHandler(tt.args.ctx, tt.args.router)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildTCPHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTCPHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
