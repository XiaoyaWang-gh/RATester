package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/server"
)

func TestSetupServer_1(t *testing.T) {
	type args struct {
		staticConfiguration *static.Configuration
	}
	tests := []struct {
		name    string
		args    args
		want    *server.Server
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

			got, err := setupServer(tt.args.staticConfiguration)
			if (err != nil) != tt.wantErr {
				t.Errorf("setupServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
