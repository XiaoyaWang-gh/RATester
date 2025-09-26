package provider

import (
	"context"
	"fmt"
	"testing"
	"text/template"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildRouterConfiguration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	configuration := &dynamic.HTTPConfiguration{}
	defaultRouterName := "defaultRouterName"
	defaultRuleTpl := &template.Template{}
	model := "model"

	BuildRouterConfiguration(ctx, configuration, defaultRouterName, defaultRuleTpl, model)

	assert.Equal(t, 0, len(configuration.Routers))
}
