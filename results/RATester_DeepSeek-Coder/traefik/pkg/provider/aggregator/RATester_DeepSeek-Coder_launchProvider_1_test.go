package aggregator

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestLaunchProvider_1(t *testing.T) {
	type args struct {
		configurationChan chan<- dynamic.Message
		pool              *safe.Pool
		prd               provider.Provider
	}
	tests := []struct {
		name string
		args args
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

			ProviderAggregator := ProviderAggregator{}
			ProviderAggregator.launchProvider(tt.args.configurationChan, tt.args.pool, tt.args.prd)
		})
	}
}
