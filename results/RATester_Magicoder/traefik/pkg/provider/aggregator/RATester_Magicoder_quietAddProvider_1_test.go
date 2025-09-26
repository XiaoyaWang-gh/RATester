package aggregator

import (
	"fmt"
	"testing"
	"time"

	"github.com/traefik/traefik/v3/pkg/provider"
)

func TestQuietAddProvider_1(t *testing.T) {
	type fields struct {
		internalProvider          provider.Provider
		fileProvider              provider.Provider
		providers                 []provider.Provider
		providersThrottleDuration time.Duration
	}
	type args struct {
		provider provider.Provider
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			p := &ProviderAggregator{
				internalProvider:          tt.fields.internalProvider,
				fileProvider:              tt.fields.fileProvider,
				providers:                 tt.fields.providers,
				providersThrottleDuration: tt.fields.providersThrottleDuration,
			}
			p.quietAddProvider(tt.args.provider)
		})
	}
}
