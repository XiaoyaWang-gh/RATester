package router

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestBuildHTTPHandler_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		router     *runtime.RouterInfo
		routerName string
	}
	tests := []struct {
		name    string
		args    args
		want    http.Handler
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
			got, err := m.buildHTTPHandler(tt.args.ctx, tt.args.router, tt.args.routerName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.buildHTTPHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.buildHTTPHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
