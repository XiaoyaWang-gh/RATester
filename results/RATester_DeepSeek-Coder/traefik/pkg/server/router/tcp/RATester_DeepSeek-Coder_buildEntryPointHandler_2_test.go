package tcp

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestBuildEntryPointHandler_2(t *testing.T) {
	type args struct {
		ctx          context.Context
		configs      map[string]*runtime.TCPRouterInfo
		configsHTTP  map[string]*runtime.RouterInfo
		handlerHTTP  http.Handler
		handlerHTTPS http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    *Router
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
			got, err := m.buildEntryPointHandler(tt.args.ctx, tt.args.configs, tt.args.configsHTTP, tt.args.handlerHTTP, tt.args.handlerHTTPS)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.buildEntryPointHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.buildEntryPointHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
