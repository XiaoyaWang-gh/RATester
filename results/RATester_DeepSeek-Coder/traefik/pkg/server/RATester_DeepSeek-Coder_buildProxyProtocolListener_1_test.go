package server

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestBuildProxyProtocolListener_1(t *testing.T) {
	type args struct {
		ctx        context.Context
		entryPoint *static.EntryPoint
		listener   net.Listener
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

			got, err := buildProxyProtocolListener(tt.args.ctx, tt.args.entryPoint, tt.args.listener)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildProxyProtocolListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildProxyProtocolListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
