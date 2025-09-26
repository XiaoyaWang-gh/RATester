package router

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestBuildEntryPointHandler_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		entryPointName string
		configs        map[string]*runtime.RouterInfo
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
			got, err := m.buildEntryPointHandler(tt.args.ctx, tt.args.entryPointName, tt.args.configs)
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
