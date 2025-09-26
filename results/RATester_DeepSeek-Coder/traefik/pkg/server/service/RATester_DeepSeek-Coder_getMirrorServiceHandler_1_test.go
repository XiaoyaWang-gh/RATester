package service

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestGetMirrorServiceHandler_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		config *dynamic.Mirroring
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
			got, err := m.getMirrorServiceHandler(tt.args.ctx, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.getMirrorServiceHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.getMirrorServiceHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
