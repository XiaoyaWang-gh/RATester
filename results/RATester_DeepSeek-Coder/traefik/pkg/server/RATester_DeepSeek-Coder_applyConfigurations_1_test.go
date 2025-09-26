package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestApplyConfigurations_1(t *testing.T) {
	type fields struct {
		providerAggregator     provider.Provider
		defaultEntryPoints     []string
		allProvidersConfigs    chan dynamic.Message
		newConfigs             chan dynamic.Configurations
		requiredProvider       string
		configurationListeners []func(dynamic.Configuration)
		routinesPool           *safe.Pool
	}
	type args struct {
		ctx context.Context
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

			c := &ConfigurationWatcher{
				providerAggregator:     tt.fields.providerAggregator,
				defaultEntryPoints:     tt.fields.defaultEntryPoints,
				allProvidersConfigs:    tt.fields.allProvidersConfigs,
				newConfigs:             tt.fields.newConfigs,
				requiredProvider:       tt.fields.requiredProvider,
				configurationListeners: tt.fields.configurationListeners,
				routinesPool:           tt.fields.routinesPool,
			}
			c.applyConfigurations(tt.args.ctx)
		})
	}
}
