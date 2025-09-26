package server

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	gokitmetrics "github.com/go-kit/kit/metrics"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNewTCPEntryPoint_1(t *testing.T) {
	type args struct {
		ctx                  context.Context
		name                 string
		config               *static.EntryPoint
		hostResolverConfig   *types.HostResolverConfig
		openConnectionsGauge gokitmetrics.Gauge
	}
	tests := []struct {
		name    string
		args    args
		want    *TCPEntryPoint
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

			got, err := NewTCPEntryPoint(tt.args.ctx, tt.args.name, tt.args.config, tt.args.hostResolverConfig, tt.args.openConnectionsGauge)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTCPEntryPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTCPEntryPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
