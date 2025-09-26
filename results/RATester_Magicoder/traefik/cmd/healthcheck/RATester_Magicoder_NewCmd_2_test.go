package healthcheck

import (
	"fmt"
	"testing"

	"github.com/traefik/paerser/cli"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewCmd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	traefikConfiguration := &static.Configuration{
		// Initialize your configuration here
	}

	loaders := []cli.ResourceLoader{
		// Initialize your loaders here
	}

	cmd := NewCmd(traefikConfiguration, loaders)

	if cmd.Name != "healthcheck" {
		t.Errorf("Expected command name to be 'healthcheck', but got '%s'", cmd.Name)
	}

	if cmd.Description != `Calls Traefik /ping endpoint (disabled by default) to check the health of Traefik.` {
		t.Errorf("Expected command description to be 'Calls Traefik /ping endpoint (disabled by default) to check the health of Traefik.', but got '%s'", cmd.Description)
	}

	if cmd.Configuration != traefikConfiguration {
		t.Errorf("Expected command configuration to be '%v', but got '%v'", traefikConfiguration, cmd.Configuration)
	}

	if cmd.Run == nil {
		t.Error("Expected command run function to be set, but it was nil")
	}

	if len(cmd.Resources) != len(loaders) {
		t.Errorf("Expected command resources to be '%v', but got '%v'", loaders, cmd.Resources)
	}
}
