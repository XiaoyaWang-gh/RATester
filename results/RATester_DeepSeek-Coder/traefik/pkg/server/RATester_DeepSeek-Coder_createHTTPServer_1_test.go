package server

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/middlewares/requestdecorator"
)

func TestCreateHTTPServer_1(t *testing.T) {
	type args struct {
		ctx           context.Context
		ln            net.Listener
		configuration *static.EntryPoint
		withH2c       bool
		reqDecorator  *requestdecorator.RequestDecorator
	}
	tests := []struct {
		name    string
		args    args
		want    *httpServer
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

			got, err := createHTTPServer(tt.args.ctx, tt.args.ln, tt.args.configuration, tt.args.withH2c, tt.args.reqDecorator)
			if (err != nil) != tt.wantErr {
				t.Errorf("createHTTPServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createHTTPServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
