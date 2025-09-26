package ecs

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildServiceConfiguration_1(t *testing.T) {
	type args struct {
		ctx           context.Context
		instance      ecsInstance
		configuration *dynamic.HTTPConfiguration
	}
	tests := []struct {
		name    string
		args    args
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

			p := &Provider{}
			if err := p.buildServiceConfiguration(tt.args.ctx, tt.args.instance, tt.args.configuration); (err != nil) != tt.wantErr {
				t.Errorf("Provider.buildServiceConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
