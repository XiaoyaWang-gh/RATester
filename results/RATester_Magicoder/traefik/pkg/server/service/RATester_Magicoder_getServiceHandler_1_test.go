package service

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestgetServiceHandler_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		service dynamic.WRRService
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
			got, err := m.getServiceHandler(tt.args.ctx, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.getServiceHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.getServiceHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
