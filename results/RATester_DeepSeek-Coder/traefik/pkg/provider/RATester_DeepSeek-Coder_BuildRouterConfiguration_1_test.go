package provider

import (
	"context"
	"fmt"
	"testing"
	"text/template"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildRouterConfiguration_1(t *testing.T) {
	type args struct {
		ctx               context.Context
		configuration     *dynamic.HTTPConfiguration
		defaultRouterName string
		defaultRuleTpl    *template.Template
		model             interface{}
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

			BuildRouterConfiguration(tt.args.ctx, tt.args.configuration, tt.args.defaultRouterName, tt.args.defaultRuleTpl, tt.args.model)
		})
	}
}
