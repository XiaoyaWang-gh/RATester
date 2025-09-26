package server

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestBuildListener_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		name   string
		config *static.EntryPoint
	}
	tests := []struct {
		name    string
		args    args
		want    net.Listener
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

			got, err := buildListener(tt.args.ctx, tt.args.name, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
